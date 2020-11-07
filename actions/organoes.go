package actions

import (
	"fmt"
	"net/http"
	"unictelezioni/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Organo)
// DB Table: Plural (organoes)
// Resource: Plural (Organoes)
// Path: Plural (/organoes)
// View Template Folder: Plural (/templates/organoes/)

// OrganoesResource is the resource for the Organo model
type OrganoesResource struct {
	buffalo.Resource
}

// List gets all Organoes. This function is mapped to the path
// GET /organoes
func (v OrganoesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	organoes := &models.Organoes{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Organoes from the DB
	if err := q.All(organoes); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(organoes))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(organoes))
	}).Respond(c)
}

// Show gets the data for one Organo. This function is mapped to
// the path GET /organoes/{organo_id}
func (v OrganoesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Organo
	organo := &models.Organo{}

	// To find the Organo the parameter organo_id is used.
	if err := tx.Find(organo, c.Param("organo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("organo", organo)

		return c.Render(http.StatusOK, r.HTML("/organoes/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(organo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(organo))
	}).Respond(c)
}

// Create adds a Organo to the DB. This function is mapped to the
// path POST /organoes
func (v OrganoesResource) Create(c buffalo.Context) error {
	// Allocate an empty Organo
	organo := &models.Organo{}

	// Bind organo to the html form elements
	if err := c.Bind(organo); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(organo)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(organo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(organo))
	}).Respond(c)
}

// Update changes a Organo in the DB. This function is mapped to
// the path PUT /organoes/{organo_id}
func (v OrganoesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Organo
	organo := &models.Organo{}

	if err := tx.Find(organo, c.Param("organo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Organo to the html form elements
	if err := c.Bind(organo); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(organo)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(organo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(organo))
	}).Respond(c)
}

// Destroy deletes a Organo from the DB. This function is mapped
// to the path DELETE /organoes/{organo_id}
func (v OrganoesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Organo
	organo := &models.Organo{}

	// To find the Organo the parameter organo_id is used.
	if err := tx.Find(organo, c.Param("organo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(organo); err != nil {
		return err
	}

	return responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(organo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(organo))
	}).Respond(c)
}
