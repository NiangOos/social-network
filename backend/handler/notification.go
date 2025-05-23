package handler

import (
	"fmt"
	"log"
	"net/http"
	"server/lib"
	"server/models"
	"strconv"
)

func HandleGetNotifications(w http.ResponseWriter, r *http.Request) {
	lib.AddCorsGet(w, r)

	_, ok := IsAuthenticated(r)

	var apiError ApiError

	if !ok {
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	sessionToken := r.Header.Get("Authorization")
	session, err := models.SessionRepo.GetSession(sessionToken)
	if err != nil {
		apiError.Error = "Go connect first !"
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	userId, err := strconv.Atoi(session.UserID)
	if err != nil {
		apiError.Error = "Error getting user."
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	if r.Method == http.MethodGet {

		notifications, err := models.NotifRepo.GetNotificationsByUserID(userId)
		if err != nil {
			fmt.Println(" --- No notifs retrieved ! ")
			WriteJSON(w, http.StatusUnauthorized, apiError)
			return
		}
		messagesPreview, err := models.MessageRepo.GetMessagePreviewsForAnUser(userId)
		if err != nil {
			log.Println("🚀 ~ func GetMessagePreviewsForAnUser ~ err:", err)
			var apiError ApiError
			apiError.Error = "Not found Messages Previews"
			WriteJSON(w, http.StatusInternalServerError, apiError)
			return
		}
		response := make(map[string]interface{})

		response["notifications"] = notifications
		response["messagesPreview"] = messagesPreview

		lib.WriteJSONResponse(w, r, response)
	}
}

func HandleUpdateNotif(w http.ResponseWriter, r *http.Request) {
	lib.AddCorsPost(w, r)

	_, ok := IsAuthenticated(r)

	var apiError ApiError

	if !ok {
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	sessionToken := r.Header.Get("Authorization")
	session, err := models.SessionRepo.GetSession(sessionToken)
	if err != nil {
		apiError.Error = "Go connect first !"
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}
	_, err = strconv.Atoi(session.UserID)
	if err != nil {
		apiError.Error = "Error getting user."
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	notifId := r.FormValue("notifId")
	intNotifId, err := strconv.Atoi(fmt.Sprintf("%v", notifId))

	if err != nil {
		http.Error(w, "Erreur group doesn't exist", http.StatusBadRequest)
		return
	}
	err = models.NotifRepo.UpdateNotificationStatus(intNotifId, true)

	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}
}

func DeleteAllNotif(w http.ResponseWriter, r *http.Request) {
	lib.AddCorsGet(w, r)

	_, ok := IsAuthenticated(r)

	var apiError ApiError

	if !ok {
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	sessionToken := r.Header.Get("Authorization")
	session, err := models.SessionRepo.GetSession(sessionToken)
	if err != nil {
		apiError.Error = "Go connect first !"
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	userId, err := strconv.Atoi(session.UserID)
	if err != nil {
		apiError.Error = "Error getting user."
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	err = models.NotifRepo.DeleteNotificationbyUser(userId)

	if err != nil {
		apiError.Error = "Error getting user."
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}

	notifications, err := models.NotifRepo.GetNotificationsByUserID(userId)
	if err != nil {
		fmt.Println(" --- No notifs retrieved ! ")
		WriteJSON(w, http.StatusUnauthorized, apiError)
		return
	}
	response := make(map[string]interface{})

	response["notifications"] = notifications

	lib.WriteJSONResponse(w, r, response)

}
