# Create a network layer
resource "google_compute_network" "shipy-network" {
  name = "${var.platform-name}"
}

# Create a firewall with some sane defaults, allowing ports 22, 80 and 443 to be open
# This is ssh,  http and https.
resource "google_compute_firewall" "ssh" {
  name = "${var.platform-name}-ssh"
  network = "${google_compute_network.shipy-network.name}"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports = [ "22", "80", "443" ]
  }

  source_ranges = ["0.0.0.0/0"]
}

# Create a new DNS zone
resource "google_dns_managed_zone" "shipy-freight" {
  dns_name = "shipyfreight.com."
  name = "shipyfreight-com"
  description = "shipyfreight.com DNS zone"
}

# Create a new subnet for our platform within our selected region
resource "google_compute_subnetwork" "shipy-freight" {
  ip_cidr_range = "10.1.2.0/24"
  name = "dev-${var.platform-name}-${var.gcloud-region}"
  network = "${google_compute_network.shipy-network.self_link}"
  region = "${var.gcloud-region}"
}

# Create a container cluster called 'shipy-freight-cluster'
# Attach new cluster to our network and our subnet,
# Ensures at least one instance is running
resource "google_container_cluster" "shipy-freight-cluster" {
  name = "shipy-freight-cluster"
  network = "${google_compute_network.shipy-network.name}"
  subnetwork = "${google_compute_subnetwork.shipy-freight.name}"
  zone = "${var.gcloud-zone}"

  initial_node_count = 1
  master_auth {
      username = ""
      password = ""
  }

  node_config {
    # Define the type/size instance to use
    # Standard is a sensible starting point
    machine_type = "n1-standard-1"

    # Grant OAuth access to the following APIs within the cluster
    oauth_scopes = [
      "https://www.googleapis.com/auth/projecthosting",
      "https://www.googleapis.com/auth/devstorage.full_control",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}

# Create a new DNS range for cluster
resource "google_dns_record_set" "dev-k8s-endpoint-shipy-freight" {
  managed_zone = "${google_dns_managed_zone.shipy-freight.name}"
  name = "k8s.dev.${google_dns_managed_zone.shipy-freight.dns_name}"
  ttl = 300
  type = "A"
  rrdatas = ["${google_container_cluster.shipy-freight-cluster.endpoint}"]
}



