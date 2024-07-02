document.addEventListener('DOMContentLoaded', () => {
    const token=jwt_decode(sessionStorage.getItem('token'));
    const employerId = token.ID 
    const jobList = document.getElementById('jobList');

    fetchJobs(employerId);

    async function fetchJobs(employerId) {
        try {
            const response = await fetch(`http://127.0.0.1:8081/getJob/${employerId}`);
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

            const removeBtn = document.createElement('button');
            removeBtn.className = 'remove-btn';
            removeBtn.textContent = 'Remove Job';
            removeBtn.addEventListener('click', () => {
                if (confirm('Are you sure you want to remove this job?')) {
                    removeJob(job.ID, jobDiv);
                }
            });
            jobDiv.appendChild(removeBtn);

            jobList.appendChild(jobDiv);
        });
    }

    async function removeJob(jobId, jobDiv) {
        try {
            const response = await fetch(`http://127.0.0.1:8081/deleteJob/${jobId}`, {
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
