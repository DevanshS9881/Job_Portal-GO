// script.js

document.addEventListener("DOMContentLoaded", () => {
    const URLParams=new URLSearchParams(window.location.search)
    const token=sessionStorage.getItem('token')
    const decoded=jwt_decode(token)
    const employerId = decoded.ID; // Replace with the actual employer ID
    const jobId = URLParams.get('ji'); // Replace with the actual job ID
    const endpoint = `http://127.0.0.1:8081/review/${employerId}/${jobId}`;

    fetch(endpoint,{
        headers: {
            'Authorization':`Bearer ${token}`,
        }
    })
        .then(response => response.json())
        .then(data => {
            if (data.applications) {
                populateApplications(data.applications);
            } else {
                console.error("No applications found");
            }
        })
        .catch(error => {
            console.error("Error fetching data:", error);
        });

    function populateApplications(applications) {
        const applicationsList = document.getElementById("applications-list");
        applicationsList.innerHTML = "";

        applications.forEach(application => {
            const row = document.createElement("tr");

            const nameCell = document.createElement("td");
            nameCell.textContent = `Applicant ${application.EmployeeID}`;
            row.appendChild(nameCell);

            const dateCell = document.createElement("td");
            dateCell.textContent = new Date(application.CreatedAt).toLocaleDateString();
            row.appendChild(dateCell);

            const statusCell = document.createElement("td");
            statusCell.textContent = application.Review || "Pending";
            row.appendChild(statusCell);

            const viewCell = document.createElement("td");
            const viewButton = document.createElement("button");
            viewButton.textContent = "View Details";
            viewButton.addEventListener("click", () => {
                viewDetails(application);
            });
            viewCell.appendChild(viewButton);
            row.appendChild(viewCell);

            applicationsList.appendChild(row);
        });
    }

    function viewDetails(application) {
        alert(`Details for Applicant ${application.EmployeeID}:\n\n${application.Letter}`);
    }
});
