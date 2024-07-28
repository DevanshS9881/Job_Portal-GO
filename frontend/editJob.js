document.addEventListener('DOMContentLoaded', () => {
    const jobId = new URLSearchParams(window.location.search).get('ji');
    const token=sessionStorage.getItem('token')
    const dtoken=jwt_decode(token);
    const id = dtoken.ID ;
    if (!jobId) {
        alert('Job ID is missing');
        return;
    }

    fetchJobDetails(jobId);

    async function fetchJobDetails(jobId) {
        try {
            const response = await fetch(`http://127.0.0.1:8081/getJob/${jobId}`, {
                method: 'GET',
                headers:{
                    'Authorization':`Bearer ${sessionStorage.getItem('token')}`
                }
                });

            const job = await response.json();
            if (job.data) {
                populateForm(job.data);
            } else {
                alert('Failed to fetch job details');
            }
        } catch (error) {
            console.error('Error fetching job details:', error);
        }
    }

    function populateForm(job) {
        document.getElementById('Profile').value = job.Profile || '';
        document.getElementById('Company').value = job.Comapny || '';
        document.getElementById('Experience').value = job.Experience || '';
        document.getElementById('Qualify').value = job.Qualification || '';
        document.getElementById('Location').value = job.Location || '';
        document.getElementById('Salary').value = job.Salary || '';
        document.getElementById('Desc').value = job.Desc || '';
        
        if (job.Status === 'Available') {
            document.getElementById('AvaliableStatus').checked = true;
        } else {
            document.getElementById('UnavaliableStatus').checked = true;
        }
    }

    document.getElementById('submitUpdate').addEventListener('click', async () => {
        const profile = document.getElementById('Profile').value;
        const company = document.getElementById('Company').value;
        const experience = document.getElementById('Experience').value;
        const qualify = document.getElementById('Qualify').value;
        const location = document.getElementById('Location').value;
        const salary = document.getElementById('Salary').value;
        const desc = document.getElementById('Desc').value;
        const status = document.querySelector('input[name="Status"]:checked').value;

        try {
            const response = await fetch(`http://127.0.0.1:8081/updateJob/${jobId}/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${sessionStorage.getItem('token')}`
                },
                body: JSON.stringify({
                    Profile: profile,
                    Company: company,
                    Experience: experience,
                    Qualify: qualify,
                    Location: location,
                    Salary: salary,
                    Desc: desc,
                    Status: status
                })
            });

            if (response.ok) {
                alert('Job updated successfully');
                window.location.href = 'posted.html'; // Redirect to job list or another page
            } else {
                alert('Failed to update job');
            }
        } catch (error) {
            console.error('Error updating job:', error);
        }
    });
});
