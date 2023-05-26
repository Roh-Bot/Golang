package Controllers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SessionMiddleware(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), "session-name")

			// Retrieve user data from the session, if available
			if userID, ok := session.Values["userID"].(int); ok {
				// User is logged in, perform necessary actions
				// For example, set the user information in the context
				c.Set("userID", userID)
			}

			// Call the next handler in the chain
			if err := next(c); err != nil {
				c.Error(err)
			}

			// Save the session after handling the request
			session.Save(c.Request(), c.Response())

			return nil
		}
	}
}

func LoginHandler(c echo.Context, store sessions.Store) error {
	// Perform login validation and retrieve the user's ID
	userID := 123

	// Create a new session and set the user's ID in the session
	session, _ := store.Get(c.Request(), "session-name")
	session.Values["userID"] = userID
	session.Save(c.Request(), c.Response())

	// Redirect or respond with success message
	return c.JSON(http.StatusOK, "Logged in successfully")
}

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("userID")
		if userID == nil {
			// User is not logged in, redirect or respond with an error
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		// User is logged in, call the next handler in the chain
		return next(c)
	}
}
func ProtectedHandler(c echo.Context) error {
	userID := c.Get("userID").(int)
	// Use the userID to perform actions for the logged-in user
	// For example, fetch user data from the database or perform other operations

	// Return a response or render a template
	return c.JSON(http.StatusOK, fmt.Sprintf("Protected route accessed by user %d", userID))
}
