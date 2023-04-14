// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LinkMediaForm struct {
	// To send media through form data, set `Content-Type` as `multipart/form-data`.
	//
	// This request has three possible components:
	// * A media key name, any alphanumeric string (e.g. `profile_picture` in the following code snippet).
	// * A value, indicating the path to the media (e.g. the paths specified after `@` in the following code snippet).
	// * Optional media metadata. Stringified JSON data, sent as a value for the media key name.
	//
	//  Here's an example request. In the path, replace `<OBJECT>` with whatever endpoint you want to reach, e.g. `entities`, `alerts`, etc.
	//
	// ```shell
	// curl -X PUT \
	// 'https://{url}/v1/<OBJECT>/{unit21_id}/link-media' \
	// -H 'Content-Type: multipart/form-data' \
	// -H 'u21-key: <YOUR_API_KEY>' \
	// --form 'profile_picture=@/src/103031/images/profile_picture.jpg' \
	// --form 'profile_picture={"media_type": "IMAGE_FACE_IMAGE", "source": "iPhone_selfie", "timestamp": 1572673229}'
	// --form 'document_front=@/src/103031/images/document_front.jpg' \
	// --form 'document_front={"media_type": "IMAGE_ID_CARD_FRONT", "source": "passport_app", "timestamp": 1572673229}'
	// ```
	//
	FormData interface{} `json:"form_data,omitempty"`
}
