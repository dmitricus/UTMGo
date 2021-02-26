package controllers

import (
	_ "html/template"
)

//func RequireLogin(h http.Handler, m *model.Model) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if u := context.Get(r, "user"); u != nil {
//			h.ServeHTTP(w, r)
//		} else {
//			http.Redirect(w, r, "/login", 302)
//		}
//	})
//}
//
//func requireAdmin(h http.Handler, m *model.Model) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		session, err := store.Get(r, "session")
//		if err != nil {
//			log.Printf("requireAdmin: err: %s\n", err)
//			return
//		}
//		r = context.Set(r, "session", session)
//		if isAdmin, ok := session.Values["is_admin"]; ok {
//			if isAdmin == false {
//				http.Error(w, "Admin required", http.StatusNotFound)
//				return
//			}
//		} else {
//			http.Redirect(w, r, "/login", 302)
//		}
//		h.ServeHTTP(w, r)
//	}
//}
