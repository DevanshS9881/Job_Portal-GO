document.addEventListener('DOMContentLoaded',function() {
    const token = sessionStorage.getItem('token');
    console.log(token);

    if (token) {
        const decoded = jwt_decode(token);
        console.log(decoded);
        const id = decoded.ID;
        console.log(id);

        const endpointUrl = `http://127.0.0.1:8082/addJob/${id}`;
        const role = decoded.role;
        console.log(role);
        document.getElementById('submitUpdate').addEventListener('click',function(){
            createJob(role,endpointUrl,token)
        })
    }
    else{
        alert("Please login");
        window.location.href = 'http://127.0.0.1:3004/index5.html';
    }

    
});

function createJob(role,endpointUrl,token){
    const jobData={
        Profile: document.getElementById('Profile').value,
        Comapny: document.getElementById('Company').value,
        Experience: document.getElementById('Experience').value,
        Qualification: document.getElementById('Qualify').value,
        Location: document.getElementById('Location').value,
        Salary: document.getElementById('Salary').value,
        Desc: document.getElementById('Desc').value,
        Status: document.querySelector('input[name="Status"]:checked').value,
    }
    fetch(endpointUrl,{
        method:'POST',
        headers:{
            'Content-Type':'application/json',
            'Authorization':`Bearer ${token}`
        },
        body:JSON.stringify(jobData),
    })
    .then(response => response.json())
    .then(data => {
        console.log("Job Added : ",data);
        alert("Job Added Successfully");
        window.location.reload();
    })
    .catch(error => {
        console.log("Error adding data",error);
        alert("Error adding data",error);
    });
}