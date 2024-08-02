document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const jobId = urlParams.get('jobID');
    const token = jwt_decode(sessionStorage.getItem('token'));
    const id = token.ID;

    document.getElementById('applyBt').addEventListener('click', async function() {
        // event.preventDefault();
        console.log("Yes");
        const applyLetter = document.getElementById('letter').value;

        const applicationData = {
            JobsId: parseInt(jobId, 10),
            Letter: applyLetter,
        };

        try {
            const response = await fetch(`https://code-backend-backend.onrender.com/apply/${id}/${jobId}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${sessionStorage.getItem('token')}`, // Corrected syntax here
                },
                body: JSON.stringify(applicationData),
            });

            const result = await response.json();
            if (response.ok) {
                alert('Application submitted successfully!');
                window.location.href="apply.html"
            } else {
                if(result.Message=="Invalid Employee" || result.Message=="Invalid Employer")
                    alert("Please Update Your Profile in Profile Section");
                if(result.Message=="You have already applied for this job")
                    alert("You have already applied for this job");
                console.error('Failed to submit application:', result.Message);
            }
        } catch (error) {
            console.error('Error submitting application:', error);
        }

    });
}); 