package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cycloidio/sqlr"
	"github.com/xescugc/chaoswall/hold"
	"github.com/xescugc/chaoswall/wall"
	"golang.org/x/xerrors"
)

type WallRepository struct {
	querier sqlr.Querier
}

type dbWall struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Canonical sql.NullString
	Image     []byte
}

func newDbWall(w wall.Wall) *dbWall {
	return &dbWall{
		ID:        ToNullInt64(w.ID),
		Name:      ToNullString(w.Name),
		Canonical: ToNullString(w.Canonical),
		Image:     w.Image,
	}
}

func (w *dbWall) toDomainEntity() *wall.Wall {
	return &wall.Wall{
		ID:        ToUint32(w.ID),
		Name:      w.Name.String,
		Canonical: w.Canonical.String,
		Image:     w.Image,
	}
}

func (w *dbWall) scanFields() []interface{} {
	return []interface{}{
		&w.ID,
		&w.Name,
		&w.Canonical,
		&w.Image,
	}
}

func dbWallFields(prefix string) string {
	fields := []string{"id", "name", "canonical", "image"}
	s := ""
	for i, f := range fields {
		format := " %s.%s,"
		if i == len(fields)-1 {
			format = " %s.%s"
		}

		s += fmt.Sprintf(format, prefix, f)
	}
	return s
}

// NewWallRepository returns an implementation of the wall.Repository
func NewWallRepository(db sqlr.Querier) wall.Repository {
	return &WallRepository{
		querier: db,
	}
}

func (r *WallRepository) Create(ctx context.Context, gCan string, w wall.Wall) (uint32, error) {
	dbw := newDbWall(w)
	res, err := r.querier.ExecContext(ctx, `
		INSERT INTO walls (name, canonical, image, gym_id)
		SELECT ?, ?, ?, g.id
		FROM
			(
				SELECT gyms.id
				FROM gyms
				WHERE gyms.canonical = ?
			) AS g
	`, dbw.Name, dbw.Canonical, dbw.Image, gCan)

	if err != nil {
		return 0, xerrors.Errorf("failed to ExecContext: %w", err)
	}

	id, err := LastInsertedID(res)
	if err != nil {
		return 0, xerrors.Errorf("failed to get LastInsertedID: %w", err)
	}

	return id, nil
}

func (r *WallRepository) FilterWithHolds(ctx context.Context, gCan string) ([]*wall.WithHolds, error) {
	rows, err := r.querier.QueryContext(ctx, fmt.Sprintf(`
		SELECT %s, %s
		FROM walls AS w
		JOIN gyms AS g
			ON g.id = w.gym_id
		JOIN holds AS h
			ON h.wall_id = w.id
		WHERE g.canonical = ?
		ORDER BY w.id
		`, dbWallFields("w"), dbHoldFields("h")), gCan,
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	walls, err := scanWallsWithHolds(rows)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanWallsWithHolds: %w", err)
	}
	return walls, nil
}

func (r *WallRepository) FindByCanonicalWithHolds(ctx context.Context, gCan, wCan string) (*wall.WithHolds, error) {
	rows, err := r.querier.QueryContext(ctx, fmt.Sprintf(`
		SELECT %s, %s
		FROM walls AS w
		JOIN gyms AS g
			ON g.id = w.gym_id
		JOIN holds AS h
			ON h.wall_id = w.id
		WHERE g.canonical = ? AND w.canonical = ?
	`, dbWallFields("w"), dbHoldFields("h")), gCan, wCan)
	if err != nil {
		return nil, xerrors.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	ws, err := scanWallsWithHolds(rows)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanWallsWithHolds: %w", err)
	}

	return ws[0], nil
}

func (r *WallRepository) UpdateByCanonical(ctx context.Context, gCan, wCan string, w wall.Wall) error {
	dbw := newDbWall(w)
	_, err := r.querier.ExecContext(ctx, `
		UPDATE walls AS	w
		JOIN gyms AS g
			ON g.id = w.gym_id
		SET w.name = ?, w.canonical = ?
		WHERE g.canonical = ? AND w.canonical = ?
	`, dbw.Name, dbw.Canonical, gCan, wCan)

	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

func (r *WallRepository) DeleteByCanonical(ctx context.Context, gCan, wCan string) error {
	_, err := r.querier.ExecContext(ctx, `
		DELETE w
		FROM walls AS w
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ?
	`, gCan, wCan)
	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

// scanWallWithHolds scans and returns a Wall from a row
func scanWallWithHolds(s sqlr.Scanner) (*wall.WithHolds, error) {
	var w dbWall
	var h dbHold

	err := s.Scan(joinFields(w.scanFields(), h.scanFields())...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("wall not found")
		}
		return nil, xerrors.Errorf("failed to Scan: %w", err)
	}

	return &wall.WithHolds{
		Wall:  *w.toDomainEntity(),
		Holds: []hold.Hold{*h.toDomainEntity()},
	}, nil
}

// scanWallsWithHolds scans and returns Walls from the rows,
// the caller is the one in charge of closing the rows
func scanWallsWithHolds(rows *sql.Rows) ([]*wall.WithHolds, error) {
	var (
		ws   []*wall.WithHolds
		idxs = make(map[uint32]int)
	)

	for rows.Next() {
		w, err := scanWallWithHolds(rows)
		if err != nil {
			return nil, xerrors.Errorf("failed to scanWall: %w", err)
		}

		if idx, ok := idxs[w.ID]; ok {
			ws[idx].Holds = append(ws[idx].Holds, w.Holds...)
		} else {
			ws = append(ws, w)
			idxs[w.ID] = len(ws) - 1
		}
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("failed to scan rows: %w", err)
	}

	return ws, nil
}
