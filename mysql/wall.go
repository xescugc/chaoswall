package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/xescugc/chaoswall/wall"
	"golang.org/x/xerrors"
)

type WallRepository struct {
	querier Querier
}

type dbWall struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Canonical sql.NullString
}

func newDbWall(w wall.Wall) *dbWall {
	return &dbWall{
		ID:        ToNullInt64(w.ID),
		Name:      ToNullString(w.Name),
		Canonical: ToNullString(w.Canonical),
	}
}

func (w *dbWall) toDomainEntity() *wall.Wall {
	return &wall.Wall{
		ID:        ToUint32(w.ID),
		Name:      w.Name.String,
		Canonical: w.Canonical.String,
	}
}

func (w *dbWall) scanFields() []interface{} {
	return []interface{}{
		&w.ID,
		&w.Name,
		&w.Canonical,
	}
}

func dbWallFields(prefix string) string {
	fields := []string{"id", "name", "canonical"}
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
func NewWallRepository(db Querier) wall.Repository {
	return &WallRepository{
		querier: db,
	}
}

func (r *WallRepository) Create(ctx context.Context, gCan string, w wall.Wall) (uint32, error) {
	dbw := newDbWall(w)
	res, err := r.querier.ExecContext(ctx, `
		INSERT INTO walls (name, canonical, gym_id)
		SELECT ?, ?, g.id
		FROM
			(
				SELECT gyms.id
				FROM gyms
				WHERE gyms.canonical = ?
			) AS g
	`, dbw.Name, dbw.Canonical, gCan)

	if err != nil {
		return 0, xerrors.Errorf("failed to ExecContext: %w", err)
	}

	id, err := LastInsertedID(res)
	if err != nil {
		return 0, xerrors.Errorf("failed to get LastInsertedID: %w", err)
	}

	return id, nil
}

func (r *WallRepository) Filter(ctx context.Context, gCan string) ([]*wall.Wall, error) {
	rows, err := r.querier.QueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM walls AS w
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ?
		ORDER BY w.id
		`, dbWallFields("w")), gCan,
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	walls, err := scanWalls(rows)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanWalls: %w", err)
	}
	return walls, nil
}

func (r *WallRepository) FindByCanonical(ctx context.Context, gCan, wCan string) (*wall.Wall, error) {
	row := r.querier.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM walls AS w
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ?
	`, dbWallFields("w")), gCan, wCan)

	w, err := scanWall(row)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanWall: %w", err)
	}

	return w, nil
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

// scanWall scans and returns a Wall from a row
func scanWall(s Scanner) (*wall.Wall, error) {
	var g dbWall

	err := s.Scan(g.scanFields()...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failed to Scan: %w", err)
	}

	return g.toDomainEntity(), nil
}

// scanWalls scans and returns Walls from the rows,
// the caller is the one in charge of closing the rows
func scanWalls(rows *sql.Rows) ([]*wall.Wall, error) {
	var gs []*wall.Wall

	for rows.Next() {
		g, err := scanWall(rows)
		if err != nil {
			return nil, xerrors.Errorf("failed to scanWall: %w", err)
		}
		gs = append(gs, g)
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("failed to scan rows: %w", err)
	}

	return gs, nil
}
