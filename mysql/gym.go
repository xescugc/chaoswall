package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cycloidio/sqlr"
	"github.com/xescugc/chaoswall/gym"
	"golang.org/x/xerrors"
)

type GymRepository struct {
	querier sqlr.Querier
}

type dbGym struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Canonical sql.NullString
}

func newDbGym(g gym.Gym) *dbGym {
	return &dbGym{
		ID:        ToNullInt64(g.ID),
		Name:      ToNullString(g.Name),
		Canonical: ToNullString(g.Canonical),
	}
}

func (g *dbGym) toDomainEntity() *gym.Gym {
	return &gym.Gym{
		ID:        ToUint32(g.ID),
		Name:      g.Name.String,
		Canonical: g.Canonical.String,
	}
}

func (g *dbGym) scanFields() []interface{} {
	return []interface{}{
		&g.ID,
		&g.Name,
		&g.Canonical,
	}
}

func dbGymFields(prefix string) string {
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

// NewGymRepository returns an implementation of the gym.Repository
func NewGymRepository(db sqlr.Querier) gym.Repository {
	return &GymRepository{
		querier: db,
	}
}

func (r *GymRepository) Create(ctx context.Context, g gym.Gym) (uint32, error) {
	res, err := r.querier.ExecContext(ctx, `
		INSERT INTO gyms(name, canonical)
		VALUES (?, ?)
	`, g.Name, g.Canonical)

	if err != nil {
		return 0, xerrors.Errorf("failed to ExecContext: %w", err)
	}

	id, err := LastInsertedID(res)
	if err != nil {
		return 0, xerrors.Errorf("failed to get LastInsertedID: %w", err)
	}

	return id, nil
}

func (r *GymRepository) Filter(ctx context.Context) ([]*gym.Gym, error) {
	rows, err := r.querier.QueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM gyms AS g
		ORDER BY g.id
		`, dbGymFields("g")),
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	gyms, err := scanGyms(rows)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanGyms: %w", err)
	}
	return gyms, nil
}

func (r *GymRepository) FindByCanonical(ctx context.Context, gCan string) (*gym.Gym, error) {
	row := r.querier.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM gyms AS g
		WHERE g.canonical = ?
	`, dbGymFields("g")), gCan)

	g, err := scanGym(row)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanGym: %w", err)
	}

	return g, nil
}

func (r *GymRepository) UpdateByCanonical(ctx context.Context, gCan string, g gym.Gym) error {
	dbg := newDbGym(g)
	_, err := r.querier.ExecContext(ctx, `
		UPDATE gyms AS g
		SET g.name = ?, g.canonical = ?
		WHERE g.canonical = ?
	`, dbg.Name, dbg.Canonical, gCan)

	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

func (r *GymRepository) DeleteByCanonical(ctx context.Context, gCan string) error {
	_, err := r.querier.ExecContext(ctx, `
		DELETE FROM gyms
		WHERE canonical=?
	`, gCan)
	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

// scanGym scans and returns a Gym from a row
func scanGym(s sqlr.Scanner) (*gym.Gym, error) {
	var g dbGym

	err := s.Scan(g.scanFields()...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failed to Scan: %w", err)
	}

	return g.toDomainEntity(), nil
}

// scanGyms scans and returns Gyms from the rows,
// the caller is the one in charge of closing the rows
func scanGyms(rows *sql.Rows) ([]*gym.Gym, error) {
	var gs []*gym.Gym

	for rows.Next() {
		g, err := scanGym(rows)
		if err != nil {
			return nil, xerrors.Errorf("failed to scanGym: %w", err)
		}
		gs = append(gs, g)
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("failed to scan rows: %w", err)
	}

	return gs, nil
}
