document.addEventListener('DOMContentLoaded', function() {
    const token = sessionStorage.getItem('token');
    console.log(token);

    if (token) {
        const decoded = jwt_decode(token);
        console.log(decoded);
        const id = decoded.ID;
        console.log(id);

        const endpointUrl = `http://127.0.0.1:8081/getProfile/${id}`;
        const role = decoded.Role;

        fetchProfile(endpointUrl, token, role);
    } else {
        alert("Please login");
        window.location.href = 'http://127.0.0.1:3000/frontend/index5.html';
    }
});

function fetchProfile(endpointUrl, token, role) {
    fetch(endpointUrl, {
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(async response => {
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message);
        }
        return response.json();
    })
    .then(Data => {
        console.log(Data);
        if (role === "Employee") {
            document.getElementById('heading').textContent = "EMPLOYEE PROFILE";
            document.getElementById('name').textContent = Data.data.Name;
            document.getElementById('userId').textContent = Data.data.ID;
            document.getElementById('dob').textContent = Data.data.Employee.BirthDate;
            document.getElementById('email').textContent = Data.data.Email;
            document.getElementById('location').textContent = Data.data.Employee.City;
        } else {
            document.getElementById('heading').textContent = "EMPLOYER PROFILE";
            document.getElementById('name').textContent = Data.data.Name;
            document.getElementById('userId').textContent = Data.data.ID;
            document.getElementById('dob').textContent = Data.data.Employer.BirthDate;
            document.getElementById('email').textContent = Data.data.Email;
            document.getElementById('location').textContent = Data.data.Employer.City;
        }
    })
    .catch(error => {
        console.error('Error fetching data:', error);
    });
}
