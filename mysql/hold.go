package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cycloidio/sqlr"
	"github.com/xescugc/chaoswall/hold"
	"golang.org/x/xerrors"
)

type HoldRepository struct {
	querier sqlr.Querier
}

type dbHold struct {
	ID   sql.NullInt64
	X    sql.NullInt64
	Y    sql.NullInt64
	Size sql.NullInt64
}

func newDbHold(h hold.Hold) *dbHold {
	return &dbHold{
		ID:   ToNullInt64(h.ID),
		X:    ToNullIInt64(h.X),
		Y:    ToNullIInt64(h.Y),
		Size: ToNullIInt64(h.Size),
	}
}

func (h *dbHold) toDomainEntity() *hold.Hold {
	return &hold.Hold{
		ID:   ToUint32(h.ID),
		X:    int(h.X.Int64),
		Y:    int(h.Y.Int64),
		Size: int(h.Size.Int64),
	}
}

func (h *dbHold) scanFields() []interface{} {
	return []interface{}{
		&h.ID,
		&h.X,
		&h.Y,
		&h.Size,
	}
}

func dbHoldFields(prefix string) string {
	fields := []string{"id", "x", "y", "size"}
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

// NewHoldRepository returns an implementation of the hold.Repository
func NewHoldRepository(db sqlr.Querier) hold.Repository {
	return &HoldRepository{
		querier: db,
	}
}

func (r *HoldRepository) Create(ctx context.Context, gCan, wCan string, h hold.Hold) (uint32, error) {
	dbh := newDbHold(h)
	res, err := r.querier.ExecContext(ctx, `
		INSERT INTO holds (x, y, size, wall_id)
		SELECT ?, ?, ?, w.id
		FROM
			(
				SELECT walls.id
				FROM walls
				JOIN gyms
					ON gyms.id = walls.gym_id
				WHERE gyms.canonical = ? AND walls.canonical = ?
			) AS w
	`, dbh.X, dbh.Y, dbh.Size, gCan, wCan)

	if err != nil {
		return 0, xerrors.Errorf("failed to ExecContext: %w", err)
	}

	id, err := LastInsertedID(res)
	if err != nil {
		return 0, xerrors.Errorf("failed to get LastInsertedID: %w", err)
	}

	return id, nil
}

func (r *HoldRepository) DeleteByWallCanonical(ctx context.Context, gCan, wCan string) error {
	_, err := r.querier.ExecContext(ctx, `
		DELETE h
		FROM holds AS h
		JOIN walls AS w
			ON w.id = h.wall_id
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ?
	`, gCan, wCan)
	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

// scanHold scans and returns a Hold from a row
func scanHold(s sqlr.Scanner) (*hold.Hold, error) {
	var g dbHold

	err := s.Scan(g.scanFields()...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failed to Scan: %w", err)
	}

	return g.toDomainEntity(), nil
}

// scanHolds scans and returns Holds from the rows,
// the caller is the one in charge of closing the rows
func scanHolds(rows *sql.Rows) ([]*hold.Hold, error) {
	var gs []*hold.Hold

	for rows.Next() {
		g, err := scanHold(rows)
		if err != nil {
			return nil, xerrors.Errorf("failed to scanHold: %w", err)
		}
		gs = append(gs, g)
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("failed to scan rows: %w", err)
	}

	return gs, nil
}
