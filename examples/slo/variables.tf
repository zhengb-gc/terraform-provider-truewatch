variable "region" {
  description = "The region of TrueWatch Cloud."
  type        = string
  default     = "singapore"
}

variable "sli_uuids" {
  description = "SLI UUIDs."
  type        = list(string)
  default     = ["rul-aaaaaa", "rul-bbbbbb"]
}

variable "alert_policy_uuids" {
  description = "Alert policy UUIDs."
  type        = list(string)
  default     = ["altpl-xxxxxx"]
}

variable "tags" {
  description = "Tags for the SLO."
  type        = list(string)
  default     = ["example", "terraform"]
}
