document.addEventListener('DOMContentLoaded', () => {
    const urlParams = new URLSearchParams(window.location.search);
    const jobId = urlParams.get('jobID');
    const token =jwt_decode(sessionStorage.getItem('token'));
    const id=token.ID;

    document.getElementById('applyBt').addEventListener('click', async function() {
        // event.preventDefault();
        console.log("Yes");
        const applyLetter = document.getElementById('letter').value;

        const applicationData = {
            JobsId: parseInt(jobId,10),
            Letter: applyLetter,
        };

        try {
            const response = await fetch(`http://127.0.0.1:8081/apply/${id}/${jobId}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
                },
                body: JSON.stringify(applicationData),
            });

            const result = await response.json();
            if (response.ok) {
                alert('Application submitted successfully!');
            } else {
                console.error('Failed to submit application:', result.message);
            }
        } catch (error) {
            console.error('Error submitting application:', error);
        }

    });
});