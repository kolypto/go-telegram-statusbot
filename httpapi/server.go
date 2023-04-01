package httpapi

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-faster/errors"
	"github.com/kolypto/go-telegram-statusbot/telegram"
)

func StartServer(ctx context.Context, bind string, client telegram.TelegramClient) error {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(10 * time.Second))

	router.Post("/set-status/{status}", func(w http.ResponseWriter, r *http.Request) {
		// Get status as a known name
		status := chi.URLParam(r, "status")
		documentId, err := getEmojiDocumentId(status)
		if err != nil {
			log.Printf("Failed to parse status: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Set status
		log.Printf("setting status: %v (%v)", documentId, status)
		err = client.SetEmojiStatus(r.Context(), documentId)

		// Handle errors
		if err != nil {
			log.Printf("Failed to set status: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		} else {
			w.Write([]byte("ok"))
		}
	})

	server := &http.Server{
		Addr:    bind,
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	log.Printf("Server listening on: %v", bind)
	return server.ListenAndServe()
}

// Parse "status" as either icon name or int DocumentID
func getEmojiDocumentId(status string) (int64, error) {
	// Try to get it from the known statuses map
	documentId, ok := emojiStatuses[status]
	if ok {
		return documentId, nil
	}

	// If not known, attempt parsing it as an integer
	n, err := fmt.Sscan(status, &documentId)
	fmt.Printf("Parsing %v: %q n=%v %v\n", status, documentId, n, err)
	if err != nil {
		return documentId, errors.Wrapf(err, "failed to parse icon DocumentId")
	} else {
		return documentId, nil
	}
}

// Known statuses and their DocumentIDs
var emojiStatuses = map[string]int64{
	"icq-online":  5276100343673924208,
	"icq-privacy": 5276336566875205182,
	"icq-offline": 5278693043631825676,
	"icq-NA":      5276011390606258868,
	"icq-DND":     5276019538159219467,
	"icq-away":    5276284082374848369,
	"house":       5276303465562256118,
	"car":         5278622017757654567,
	"coffee":      5278219884264696479,
	"pedestrian":  5276300811272465264,
	"headphones":  5276308662472682873,
	"dnd":         5278531947998488827,
	"toilet":      5278385622757681011,
}
