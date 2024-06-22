document.addEventListener('DOMContentLoaded', function() {
    // Replace the URL with your actual endpoint
    const token=sessionStorage.getItem('token');
    console.log(token);
    if(token){
    const decoded=jwt_decode(token);
    console.log(decoded);
    const id=decoded.ID;
    console.log(id);
   const endpointUrl = `http://127.0.0.1:8081/getProfile/${id}`;
    fetch(endpointUrl,{
        headers:{
            'Authorization':`Bearer ${token}`
        }
    })
        .then(async response => {
            if(!response.ok){
                const errorData=await response.json();
                throw new Error(errorData.message);
            }
            return response.json();
        })
        .then(Data => {
            //console.log(Data);
            //const user=data.user;
            document.getElementById('name').textContent = Data.data.Name;
            document.getElementById('userId').textContent = Data.data.ID;
            document.getElementById('dob').textContent = Data.data.Employee.BirthDate;
            document.getElementById('email').textContent = Data.data.Email;
            document.getElementById('location').textContent = Data.data.Employee.City;
            document.getElementById('skill').textContent = Data.data.Employee.Skill;

        })
        // .catch(error => {
        //     console.error('Error fetching data:', error);
        // });
    }
    else{
        alert("Please login");
        window.location.href = 'http://127.0.0.1:3000/frontend/index5.html';
    }
});