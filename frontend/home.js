document.addEventListener('DOMContentLoaded', function () {
    const token = sessionStorage.getItem('token');
    console.log(token);
    if (token) {
        const decoded = jwt_decode(token);
        console.log(decoded);
        if (decoded.role == "Employer") {
            const postButtonHolder = document.getElementById('postButton-holder');
            const postButton = document.createElement('button');
            postButton.id = "postBt"
            postButton.type = "submit";
            postButton.textContent = "Post Job";
            postButtonHolder.appendChild(postButton);
            postButton.style.paddingLeft = "25px";
            postButton.style.paddingRight = "25px";
            postButton.style.paddingTop = "8px";
            postButton.style.paddingBottom = "8px";
            postButton.style.fontSize = "1.2rem";
            postButton.style.backgroundColor = "#407ff0"
            postButton.style.color = "whitesmoke";
            postButton.style.borderRadius = "0.5vw"
            postButton.style.margin = "2vw"
            document.getElementById('posted').textContent = "Jobs Posted";
            document.getElementById('posted').href = "posted.html";


            document.getElementById('postBt').addEventListener('click', function (event) {
                event.preventDefault();
                window.location.href = 'http://127.0.0.1:3000/frontend/postJob.html';
            })

        }
    }
    else {
      const loginLink=document.createElement('a');
      loginLink.className="link";
      loginLink.textContent="Register/Login";
      loginLink.href="index5.html";
      loginLink.style.backgroundColor="#407ff0";
      loginLink.style.padding="1vh";
      const navbar=document.querySelector('.navbar')
      navbar.appendChild(loginLink);

    }

    const profiles = [
        { id: "web-developer", profile: "Web Developer" },
        { id: "devops-manager", profile: "DevOps Manager" },
        { id: "app-developer", profile: "App Developer" },
        { id: "finance-sales", profile: "Finance & Sales" },
        { id: "human-resource", profile: "Human Resource" },
        { id: "consultancy", profile: "Consultancy" },
    ];

    profiles.forEach(({ id, profile }) => {
        var encodedProfile=encodeURIComponent(profile);
        fetch(`http://127.0.0.1:8081/jobs/profiles/${encodedProfile}`)
            .then(response => response.json())
            .then(data => {
                document.querySelector(`#${id} .jobs-count`).textContent = `${data.length} jobs`;
            })
            .catch(error => {
                document.querySelector(`#${id} .jobs-count`).textContent = "Error loading jobs";
                console.error('Error fetching jobs:', error);
            });
    });
});






//}






