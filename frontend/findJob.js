document.addEventListener('DOMContentLoaded', function() {
    const token = sessionStorage.getItem('token');
    console.log(token);
    var logoutLink = document.getElementById('logout');
        logoutLink.addEventListener('click', function(event) {
        event.preventDefault();
        var confirmLogout = confirm("Are you sure you want to logout?");
        if (confirmLogout) {
            sessionStorage.removeItem('token');
            window.location.href = 'index5.html'; 
        }
    });
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
                window.location.href = 'postJob.html';
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
    const urlParams = new URLSearchParams(window.location.search);
    const profileParam = urlParams.get('profile');

    if (profileParam) {
        document.getElementById('searchP').value = profileParam;
        fetchJobs(profileParam);
    } else {
        fetchJobs();
    }
    document.getElementById('filterBt').addEventListener('click', function() {
        const prof = document.getElementById('searchP').value;
        const location = document.getElementById('searchL').value;
        fetchJobs(prof, location);
    });

    

    // Get the modal
    var modal = document.getElementById('myModal');

    // Get the <span> element that closes the modal
    var span = document.getElementsByClassName('close')[0];

    // When the user clicks on <span> (x), close the modal
    span.onclick = function() {
        modal.style.display = 'none';
    }

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = 'none';
        }
    }
});

function fetchJobs(prof = '', location = '') {
    fetch('http://127.0.0.1:8082/allJobs') // Replace with your actual backend URL
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                displayJobs(data.data, prof, location);
            } else {
                console.error('Failed to fetch jobs:', data.message);
            }
        })
        .catch(error => console.error('Error fetching jobs:', error));
}

function displayJobs(jobs, prof, location) {
    const jobList = document.getElementById('jobList');
    jobList.innerHTML = '';
    let count = 0;

    jobs.forEach(job => {
        if ((prof === '' || job.Profile.toLowerCase().includes(prof.toLowerCase())) &&
            (location === '' || job.Location.toLowerCase().includes(location.toLowerCase()))) {

            count++;
            
            const jobBlock = document.createElement('div');
            jobBlock.className = 'job';

            const profile = document.createElement('h2');
            profile.textContent = job.Profile;
            jobBlock.appendChild(profile);

            const company = document.createElement('p');
            company.textContent = `Company: ${job.Comapny}`;
            jobBlock.appendChild(company);

            const loc = document.createElement('p');
            loc.textContent = `Location: ${job.Location}`;
            jobBlock.appendChild(loc);

            const experience = document.createElement('p');
            experience.textContent = `Experience: ${job.Experience}`;
            jobBlock.appendChild(experience);

            const qualification = document.createElement('p');
            qualification.textContent = `Qualification: ${job.Qualification}`;
            jobBlock.appendChild(qualification);

            const salary = document.createElement('p');
            salary.textContent = `Salary: ${job.Salary}`;
            jobBlock.appendChild(salary);

            const applyBtn = document.createElement('button');
            applyBtn.className = 'apply-btn';
            if(sessionStorage.getItem('token')){
                applyBtn.textContent = 'Apply for Job';
            applyBtn.addEventListener('click', () => {
                window.location.href = `apply.html?jobID=${job.ID}`;
            });
            }
            else{
                applyBtn.textContent = 'Login to Apply';
                applyBtn.addEventListener('click', () => {
                    window.location.href = `index5.html`;
                });
        }
            jobBlock.appendChild(applyBtn);


            // const desc = document.createElement('p');
            // desc.textContent = `Description: ${job.Desc}`;
            // jobBlock.appendChild(desc);

            jobBlock.addEventListener('click', function() {
                showModal(job);
            });

            jobList.appendChild(jobBlock);
        }
    });
    
    if (count == 0) {
        const jobBlock = document.createElement('div');
        jobBlock.className = 'job';

        const profile = document.createElement('h2');
        profile.textContent = "No search result found";
        profile.style.alignSelf = "center";
        jobBlock.style.border = "none";
        jobBlock.style.display = "flex";
        jobBlock.style.justifyContent = "center";
        jobBlock.style.alignItems = "center";
        jobBlock.appendChild(profile);
        jobList.appendChild(jobBlock);
    }
}

function showModal(job) {
    const token = sessionStorage.getItem('token');
    console.log(token);
    const decoded = jwt_decode(token);
    if(decoded.role=="Employer"){
        document.getElementById('applyBt').style.display='none';
    }
    const modal = document.getElementById('myModal');
    const modalJobDetails = document.getElementById('modalJobDetails');
    modalJobDetails.innerHTML = `
        <h2>${job.Profile}</h2>
        <p>Company: ${job.Comapny}</p>
        <p>Location: ${job.Location}</p>
        <p>Experience: ${job.Experience}</p>
        <p>Qualification: ${job.Qualification}</p>
        <p>Salary: ${job.Salary}</p>
        <p>Description: ${job.Desc}</p>
    `;
    modal.style.display = 'block';
}
