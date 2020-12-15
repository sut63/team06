// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/playlist"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/user"
)

// PlaylistUpdate is the builder for updating Playlist entities.
type PlaylistUpdate struct {
	config
	hooks      []Hook
	mutation   *PlaylistMutation
	predicates []predicate.Playlist
}

// Where adds a new predicate for the builder.
func (pu *PlaylistUpdate) Where(ps ...predicate.Playlist) *PlaylistUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetTitle sets the title field.
func (pu *PlaylistUpdate) SetTitle(s string) *PlaylistUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetOwnerID sets the owner edge to User by id.
func (pu *PlaylistUpdate) SetOwnerID(id int) *PlaylistUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableOwnerID(id *int) *PlaylistUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to User.
func (pu *PlaylistUpdate) SetOwner(u *User) *PlaylistUpdate {
	return pu.SetOwnerID(u.ID)
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (pu *PlaylistUpdate) AddPlaylistVideoIDs(ids ...int) *PlaylistUpdate {
	pu.mutation.AddPlaylistVideoIDs(ids...)
	return pu
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (pu *PlaylistUpdate) AddPlaylistVideos(p ...*Playlist_Video) *PlaylistUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pu *PlaylistUpdate) Mutation() *PlaylistMutation {
	return pu.mutation
}

// ClearOwner clears the owner edge to User.
func (pu *PlaylistUpdate) ClearOwner() *PlaylistUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// RemovePlaylistVideoIDs removes the playlist_videos edge to Playlist_Video by ids.
func (pu *PlaylistUpdate) RemovePlaylistVideoIDs(ids ...int) *PlaylistUpdate {
	pu.mutation.RemovePlaylistVideoIDs(ids...)
	return pu
}

// RemovePlaylistVideos removes playlist_videos edges to Playlist_Video.
func (pu *PlaylistUpdate) RemovePlaylistVideos(p ...*Playlist_Video) *PlaylistUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePlaylistVideoIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PlaylistUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Title(); ok {
		if err := playlist.TitleValidator(v); err != nil {
			return 0, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlaylistUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlaylistUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlaylistUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlaylistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlist.Table,
			Columns: playlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldTitle,
		})
	}
	if pu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.OwnerTable,
			Columns: []string{playlist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.OwnerTable,
			Columns: []string{playlist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedPlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.PlaylistVideosTable,
			Columns: []string{playlist.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.PlaylistVideosTable,
			Columns: []string{playlist.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PlaylistUpdateOne is the builder for updating a single Playlist entity.
type PlaylistUpdateOne struct {
	config
	hooks    []Hook
	mutation *PlaylistMutation
}

// SetTitle sets the title field.
func (puo *PlaylistUpdateOne) SetTitle(s string) *PlaylistUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetOwnerID sets the owner edge to User by id.
func (puo *PlaylistUpdateOne) SetOwnerID(id int) *PlaylistUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableOwnerID(id *int) *PlaylistUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to User.
func (puo *PlaylistUpdateOne) SetOwner(u *User) *PlaylistUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (puo *PlaylistUpdateOne) AddPlaylistVideoIDs(ids ...int) *PlaylistUpdateOne {
	puo.mutation.AddPlaylistVideoIDs(ids...)
	return puo
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (puo *PlaylistUpdateOne) AddPlaylistVideos(p ...*Playlist_Video) *PlaylistUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (puo *PlaylistUpdateOne) Mutation() *PlaylistMutation {
	return puo.mutation
}

// ClearOwner clears the owner edge to User.
func (puo *PlaylistUpdateOne) ClearOwner() *PlaylistUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// RemovePlaylistVideoIDs removes the playlist_videos edge to Playlist_Video by ids.
func (puo *PlaylistUpdateOne) RemovePlaylistVideoIDs(ids ...int) *PlaylistUpdateOne {
	puo.mutation.RemovePlaylistVideoIDs(ids...)
	return puo
}

// RemovePlaylistVideos removes playlist_videos edges to Playlist_Video.
func (puo *PlaylistUpdateOne) RemovePlaylistVideos(p ...*Playlist_Video) *PlaylistUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePlaylistVideoIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PlaylistUpdateOne) Save(ctx context.Context) (*Playlist, error) {
	if v, ok := puo.mutation.Title(); ok {
		if err := playlist.TitleValidator(v); err != nil {
			return nil, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}

	var (
		err  error
		node *Playlist
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlaylistUpdateOne) SaveX(ctx context.Context) *Playlist {
	pl, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pl
}

// Exec executes the query on the entity.
func (puo *PlaylistUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlaylistUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlaylistUpdateOne) sqlSave(ctx context.Context) (pl *Playlist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlist.Table,
			Columns: playlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlist.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Playlist.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playlist.FieldTitle,
		})
	}
	if puo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.OwnerTable,
			Columns: []string{playlist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.OwnerTable,
			Columns: []string{playlist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedPlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.PlaylistVideosTable,
			Columns: []string{playlist.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.PlaylistVideosTable,
			Columns: []string{playlist.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pl = &Playlist{config: puo.config}
	_spec.Assign = pl.assignValues
	_spec.ScanValues = pl.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pl, nil
}
