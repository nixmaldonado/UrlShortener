package main

const (
	EventServerStart         = "server_start"
	EventShortRequest        = "short_request"
	EventMissingSchemeOrHost = "missing_scheme_or_host"
	EventRedirectRequest     = "redirect_request"
	EventEmptyStorageFile    = "empty_storage_file"
)

const (
	ErrorParsingURL               = "error_parsing_url"
	ErrorCreatingStorage          = "error_creating_storage"
	ErrorShortRequestPayload      = "error_with_payload"
	ErrorRunningServer            = "error_running_server"
	ErrorEmptyURL                 = "error_empty_url"
	ErrorIncrementingCounter      = "error_incrementing_counter"
	ErrorFileNotExist             = "error_file_not_exist"
	ErrorFailedToReadFile         = "error_failed_to_read_file"
	ErrorUnmarshallingStorageFile = "error_unmarshalling_storage_file"
	ErrorIndentingFile            = "error_indenting_file"
	ErrorWritingToStorage         = "error_writing_to_storage"
	ErrorRenamingTempFile         = "error_renaming_temp_file"
)
