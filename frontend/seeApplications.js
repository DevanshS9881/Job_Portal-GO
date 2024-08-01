document.addEventListener("DOMContentLoaded", () => {
    const token = sessionStorage.getItem('token');
    const decoded = jwt_decode(token);
    const employeeId = decoded.ID;
    const endpoint = `http://127.0.0.1:8082/getApplications/${employeeId}`;

    fetch(endpoint, {
        headers: {
            'Authorization': `Bearer ${token}`,
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

            const jobNameCell = document.createElement("td");
            jobNameCell.textContent = application.Jobs.Profile;
            row.appendChild(jobNameCell);

            const companyCell = document.createElement("td");
            companyCell.textContent = application.Jobs.Comapny;
            row.appendChild(companyCell);

            const dateCell = document.createElement("td");
            dateCell.textContent = new Date(application.CreatedAt).toLocaleDateString();
            row.appendChild(dateCell);

           

            const statusCell = document.createElement("td");
            statusCell.textContent = application.Review || "Pending";
            row.appendChild(statusCell);

            const actionsCell = document.createElement("td");
            const viewButton = document.createElement("button");
            viewButton.textContent = "View Details";
            viewButton.addEventListener("click", () => {
                viewDetails(application);
            });
            actionsCell.appendChild(viewButton);
            row.appendChild(actionsCell);

            applicationsList.appendChild(row);
        });
    }

    function viewDetails(application) {
        const modal = document.getElementById("modal");
        const closeButton = document.querySelector(".close-button");

        document.getElementById("job-profile").textContent = `Profile: ${application.Jobs.Profile}`;
        document.getElementById("job-company").textContent = `Company: ${application.Jobs.Comapny}`;
        document.getElementById("job-location").textContent = `Location: ${application.Jobs.Location}`;
        document.getElementById("job-description").textContent = `Description: ${application.Jobs.Desc}`;
        document.getElementById("application-letter").textContent = application.Letter;

        modal.style.display = "block";

        closeButton.addEventListener("click", () => {
            modal.style.display = "none";
        });

        window.addEventListener("click", (event) => {
            if (event.target == modal) {
                modal.style.display = "none";
            }
        });

        // document.getElementById("delete-button").addEventListener("click", () => {
        //     // Handle delete action
        //     fetch(`http://127.0.0.1:8082/delete/${application.ID}`, {
        //         method: 'DELETE',
        //         headers: {
        //             'Authorization': `Bearer ${token}`,
        //         }
        //     })
        //         .then(response => {
        //             if (response.ok) {
        //                 alert("Application deleted");
        //                 modal.style.display = "none";
        //                 // Reload the page or remove the row from the table
        //                 location.reload();
        //             } else {
        //                 console.error("Failed to delete application");
        //             }
        //         })
        //         .catch(error => {
        //             console.error("Error deleting application:", error);
        //         });
        // });
    }
});
