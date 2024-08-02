document.addEventListener('DOMContentLoaded', () => {
    const token=jwt_decode(sessionStorage.getItem('token'));
    const employerId = token.ID 
    const jobList = document.getElementById('jobList');

    fetchJobs(employerId);

    async function fetchJobs(employerId) {
        try {
            const response = await fetch(`https://code-backend-backend.onrender.com/getJobs/${employerId}`);
            const jobs = await response.json();
            console.log(jobs);
            displayJobs(jobs);
        } catch (error) {
            console.error('Error fetching jobs:', error);
        }
    }

    function displayJobs(jobs) {

        jobs.data.forEach(job => {
            const jobDiv = document.createElement('div');
            jobDiv.className = 'job';

            const jobHeading = document.createElement('h3');
            jobHeading.className = 'job-heading';
            jobHeading.textContent = job.Profile;
            jobDiv.appendChild(jobHeading);

            const jobCreated = document.createElement('p');
            jobCreated.textContent = `Created on: ${new Date(job.CreatedAt).toLocaleDateString()}`;
            jobDiv.appendChild(jobCreated);

            const reviewLink=document.createElement('a');
            reviewLink.textContent=`${job.ApplicationsRecieved} X Applications`
            reviewLink.href=`review.html?ji=${job.ID}`
            jobDiv.appendChild(reviewLink)


            const removeBtn = document.createElement('button');
            removeBtn.className = 'remove-btn';
            removeBtn.textContent = 'Remove Job';
            removeBtn.style.margin='1vw'
            removeBtn.style.fontSize='100%'
            removeBtn.style.padding="0.7vw"
            removeBtn.style.borderRadius='0.5vw'
            removeBtn.addEventListener('click', () => {
                if (confirm('Are you sure you want to remove this job?')) {
                    removeJob(job.ID, jobDiv);
                }
            });
            jobDiv.appendChild(removeBtn);
            const editBtn = document.createElement('button');
            editBtn.className = 'edit-btn';
            editBtn.textContent = 'Edit Job';
            editBtn.style.margin = '1vw';
            editBtn.style.fontSize = '100%';
            editBtn.style.padding = '0.7vw';
            editBtn.style.borderRadius = '0.5vw';
            editBtn.addEventListener('click', () => {
                // Redirect to edit page with job ID as a query parameter
                window.location.href = `editJob.html?ji=${job.ID}`;
            });
            jobDiv.appendChild(editBtn);
            jobDiv.style.backgroundColor="white";
            jobDiv.style.borderRadius="1.5vw";
            jobDiv.style.padding="28px";

            jobList.appendChild(jobDiv);
        });
    }

    async function removeJob(jobId, jobDiv) {
        try {
            const response = await fetch(`https://code-backend-backend.onrender.com/deleteJob/${jobId}`, {
                method: 'DELETE',
                headers:{
                    'Authorization':`Bearer ${sessionStorage.getItem('token')}`
                }

            });
            if (response.ok) {
                jobList.removeChild(jobDiv);
            } else {
                console.error('Failed to remove job:', response.statusText);
            }
        } catch (error) {
            console.error('Error removing job:', error);
        }
    }
});
