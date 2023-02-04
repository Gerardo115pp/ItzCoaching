package handlers

import (
	"fmt"
	"io/ioutil"
	"itz_customers_service/helpers"
	"itz_customers_service/models"
	"itz_customers_service/repository"
	"itz_customers_service/server"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func ProfilePicturesHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getProfilePictureHandler(response, request)
		case http.MethodPost:
			postProfilePictureHandler(response, request)
		case http.MethodPatch:
			patchProfilePictureHandler(response, request)
		case http.MethodDelete:
			deleteProfilePictureHandler(response, request)
		case http.MethodPut:
			putProfilePictureHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putProfilePictureHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteProfilePictureHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchProfilePictureHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func postProfilePictureHandler(response http.ResponseWriter, request *http.Request) {
	var authenticated_claims map[string]any = request.Context().Value("claim_data").(map[string]any)
	var expert_id int = authenticated_claims["id"].(int)

	var public_profile *models.PublicProfile
	public_profile, err := repository.GetExpertPublicProfile(request.Context(), expert_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error getting expert public profile for if %d", expert_id))
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	file, file_header, err := request.FormFile("profile_picture")
	if err != nil {
		echo.Echo(echo.RedFG, "Error getting file from request")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if file_header.Size > (22 * 1000000) { // 22 MB
		echo.Echo(echo.RedFG, "File size is too big")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var file_bytes []byte
	is_jpeg, err := helpers.IsJPEG(file)
	if err != nil {
		echo.Echo(echo.RedFG, "Error checking if file is jpeg")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if is_jpeg {
		file_bytes, err = ioutil.ReadAll(file)
		if err != nil {
			echo.Echo(echo.RedFG, "Error reading file")
			echo.EchoErr(err)
			response.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		file_bytes, err = helpers.ConvertToJPEG(file)
		if err != nil {
			echo.Echo(echo.RedFG, "Error converting file to jpeg")
			echo.EchoErr(err)
			response.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if public_profile.Image_url == "" {
		public_profile.Image_url = fmt.Sprintf("operation_data/experts/expert_%d/profile_picture.jpg", expert_id)
		repository.UpdateExpertPublicProfile(request.Context(), public_profile)
	}

	err = ioutil.WriteFile(public_profile.Image_url, file_bytes, 0644)
	if err != nil {
		echo.Echo(echo.RedFG, "Error writing file")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusInternalServerError)

		public_profile.Image_url = ""
		repository.UpdateExpertPublicProfile(request.Context(), public_profile)

		return
	}

	response.WriteHeader(201)
	return

}

func getProfilePictureHandler(response http.ResponseWriter, request *http.Request) {
	var profile_picture_path string = strings.Replace(request.URL.Path, "/profile_pictures", ".", 1)

	file_descriptor, err := os.Open(profile_picture_path)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	defer file_descriptor.Close()

	file_header := make([]byte, 512)
	file_descriptor.Read(file_header)

	content_type := http.DetectContentType(file_header)

	file_stat, err := file_descriptor.Stat()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	var file_size string = strconv.FormatInt(file_stat.Size(), 10)

	response.Header().Set("Content-Type", content_type)
	response.Header().Set("Content-Length", file_size)

	var file_name string = strings.Split(profile_picture_path, "/")[len(strings.Split(profile_picture_path, "/"))-1]

	response.Header().Set("Content-Disposition", "attachment; filename="+file_name)

	file_descriptor.Seek(0, 0)

	var filedata []byte
	filedata, err = ioutil.ReadAll(file_descriptor)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(200)
	response.Write(filedata)
}
