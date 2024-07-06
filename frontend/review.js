document.addEventListener("DOMContentLoaded", () => {
    const URLParams = new URLSearchParams(window.location.search);
    const token = sessionStorage.getItem('token');
    const decoded = jwt_decode(token);
    const employerId = decoded.ID; // Replace with the actual employer ID
    const jobId = URLParams.get('ji'); // Replace with the actual job ID
    const endpoint = `http://127.0.0.1:8081/review/${employerId}/${jobId}`;

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

            const nameCell = document.createElement("td");
            nameCell.textContent = `Applicant ${application.Employee.Name}`;
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

            // Apply border color based on the review status
            if (application.Review === "Accepted") {
                row.style.border = "2px solid green";
            } else if (application.Review === "Rejected") {
                row.style.border = "2px solid red";
            }
            
            applicationsList.appendChild(row);
        });
    }document.addEventListener("DOMContentLoaded", () => {
        const URLParams = new URLSearchParams(window.location.search);
        const token = sessionStorage.getItem('token');
        const decoded = jwt_decode(token);
        const employerId = decoded.ID; // Replace with the actual employer ID
        const jobId = URLParams.get('ji'); // Replace with the actual job ID
        const endpoint = `http://127.0.0.1:8081/review/${employerId}/${jobId}`;
    
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
                row.classList.add("row-style");
    
                const nameCell = document.createElement("td");
                nameCell.textContent = `Applicant ${application.Employee.Name}`;
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
    
                // Apply border color and radius based on the review status
                if (application.Review === "Accepted") {
                    row.classList.add("row-accepted");
                } else if (application.Review === "Rejected") {
                    row.classList.add("row-rejected");
                }
    
                applicationsList.appendChild(row);
            });
        }
    
        function viewDetails(application) {
            const modal = document.getElementById("modal");
            const closeButton = document.querySelector(".close-button");
    
            document.getElementById("employee-name").textContent = `Name: ${application.Employee.Name}`;
            document.getElementById("employee-city").textContent = `City: ${application.Employee.City}`;
            document.getElementById("employee-birthdate").textContent = `Birth Date: ${application.Employee.Birth_Date}`;
            document.getElementById("employee-age").textContent = `Age: ${application.Employee.Age}`;
            document.getElementById("employee-bio").textContent = `Bio: ${application.Employee.Bio}`;
            document.getElementById("employee-skill").textContent = `Skill: ${application.Employee.Skill}`;
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
    
            if(application.Review){
                document.getElementById("accept-button").style.display="none";
                document.getElementById("reject-button").style.display="none";
            }
            else{
                document.getElementById("accept-button").addEventListener("click", () => {
                    // Handle accept action
                    const resp={Review: "Accepted"};
                    fetch(`http://127.0.0.1:8081/accept/${application.ID}`,{
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json' // Set the correct content type
                        },
                        body: JSON.stringify(resp),
                    })
                    .then(response => response.json())
                    .then(data => {
                        if(data.error){
                            console.error("Cannot submit response");
                        } else {
                            alert("Application accepted");
                            populateApplications([application]); // Update the row style
                        }
                    })
                    .catch(error => console.error('Error:', error));
    
                    document.getElementById("modal").style.display = "none";
                });
    
                document.getElementById("reject-button").addEventListener("click", () => {
                    // Handle reject action
                    const resp={Review: "Rejected"};
                    fetch(`http://127.0.0.1:8081/accept/${application.ID}`,{
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json' // Set the correct content type
                        },
                        body: JSON.stringify(resp),
                    })
                    .then(response => response.json())
                    .then(data => {
                        if(data.error){
                            console.error("Cannot submit response");
                        } else {
                            alert("Application rejected");
                            populateApplications([application]); // Update the row style
                        }
                    })
                    .catch(error => console.error('Error:', error));
    
                    document.getElementById("modal").style.display = "none";
                });
            }
        }
    });
    

    function viewDetails(application) {
        const modal = document.getElementById("modal");
        const closeButton = document.querySelector(".close-button");

        document.getElementById("employee-name").textContent = `Name: ${application.Employee.Name}`;
        document.getElementById("employee-city").textContent = `City: ${application.Employee.City}`;
        document.getElementById("employee-birthdate").textContent = `Birth Date: ${application.Employee.Birth_Date}`;
        document.getElementById("employee-age").textContent = `Age: ${application.Employee.Age}`;
        document.getElementById("employee-bio").textContent = `Bio: ${application.Employee.Bio}`;
        document.getElementById("employee-skill").textContent = `Skill: ${application.Employee.Skill}`;
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

        if(application.Review){
            document.getElementById("accept-button").style.display="none";
            document.getElementById("reject-button").style.display="none";
        }
        else{
            document.getElementById("accept-button").addEventListener("click", () => {
                // Handle accept action
                const resp={Review: "Accepted"};
                fetch(`http://127.0.0.1:8081/accept/${application.ID}`,{
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json' // Set the correct content type
                    },
                    body: JSON.stringify(resp),
                })
                .then(response => response.json())
                .then(data => {
                    if(data.error){
                        console.error("Cannot submit response");
                    } else {
                        alert("Application accepted");
                    }
                })
                .catch(error => console.error('Error:', error));

                document.getElementById("modal").style.display = "none";
                populateApplications([application]); // Update the row style
            });

            document.getElementById("reject-button").addEventListener("click", () => {
                // Handle reject action
                const resp={Review: "Rejected"};
                fetch(`http://127.0.0.1:8081/accept/${application.ID}`,{
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json' // Set the correct content type
                    },
                    body: JSON.stringify(resp),
                })
                .then(response => response.json())
                .then(data => {
                    if(data.error){
                        console.error("Cannot submit response");
                    } else {
                        alert("Application rejected");
                    }
                })
                .catch(error => console.error('Error:', error));

                document.getElementById("modal").style.display = "none";
                populateApplications([application]); // Update the row style
            });
        }
    }
});
