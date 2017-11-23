variable "gitea_token" {
    type = "string"
}

variable "gitea_base_url" {
    type = "string"
}

provider "gitea" {
    token = "${var.gitea_token}"
    base_url = "${var.gitea_base_url}"
}

resource "gitea_user" "bob" {
    login = "bob"
    password = "password"
    email = "bob@example.com"
    full_name = "Bob Marley"
}