variable "f5xc_ip_ranges_Americas_TCP" {
    type = list(string)
    default = ["84.54.62.0/25","185.94.142.0/25","185.94.143.0/25","159.60.190.0/24","5.182.215.0/25","84.54.61.0/25","23.158.32.0/25",]
}
variable "f5xc_ip_ranges_Americas_UDP" {
    type = list(string)
    default = ["23.158.32.0/25","84.54.62.0/25","185.94.142.0/25","185.94.143.0/25","159.60.190.0/24","5.182.215.0/25","84.54.61.0/25",]
}
variable "f5xc_ip_ranges_Europe_TCP" {
    type = list(string)
    default = ["84.54.60.0/25","185.56.154.0/25","159.60.162.0/24","159.60.188.0/24","5.182.212.0/25","5.182.214.0/25","159.60.160.0/24","5.182.213.0/25","5.182.213.128/25",]
}
variable "f5xc_ip_ranges_Europe_UDP" {
    type = list(string)
    default = ["5.182.212.0/25","185.56.154.0/25","159.60.160.0/24","5.182.213.0/25","5.182.213.128/25","5.182.214.0/25","84.54.60.0/25","159.60.162.0/24","159.60.188.0/24",]
}
variable "f5xc_ip_ranges_Asia_TCP" {
    type = list(string)
    default = ["103.135.56.0/25","103.135.56.128/25","103.135.58.128/25","159.60.189.0/24","159.60.166.0/24","103.135.57.0/25","103.135.59.0/25","103.135.58.0/25","159.60.164.0/24",]
}
variable "f5xc_ip_ranges_Asia_UDP" {
    type = list(string)
    default = ["103.135.57.0/25","103.135.56.128/25","103.135.59.0/25","103.135.58.0/25","159.60.166.0/24","159.60.164.0/24","103.135.56.0/25","103.135.58.128/25","159.60.189.0/24",]
}
