document.addEventListener('DOMContentLoaded', function() {
    const token = sessionStorage.getItem('token');
    console.log(token);

    if (token) {
        const decoded = jwt_decode(token);
        console.log(decoded);
        const id = decoded.ID;
        console.log(id);

        const endpointUrl = `http://127.0.0.1:8082/getProfile/${id}`;
        const role = decoded.role;
        console.log(role);

        fetchProfile(endpointUrl, token, role);

        // Event listener for Update Profile button
        document.getElementById('updateBt').addEventListener('click', function() {
            showUpdateForm(role);
        });

        document.getElementById('deleteBt').addEventListener('click', function() {
            const confirmDelete = confirm("Are you sure you want to delete your profile?");
            if (confirmDelete) {
                deleteProfile(id, token);
            }
        });

        // Event listener for form submission
        document.getElementById('submitUpdate').addEventListener('click', function(event) {
            event.preventDefault();
            submitUpdateForm(id, token, role);
        });
    } else {
        alert("Please login");
        window.location.href = 'index5.html';
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
            //document.getElementById('heading').textContent = "EMPLOYEE PROFILE";
            document.getElementById('name').textContent = Data.data.Name;
            document.getElementById('userId').textContent = Data.data.ID;
            document.getElementById('dob').textContent = Data.data.Employee.BirthDate;
            document.getElementById('email').textContent = Data.data.Email;
            document.getElementById('location').textContent = Data.data.Employee.City;
            document.getElementById('f6').textContent = "Skill";
            document.getElementById('f7').textContent = "Employee ID";
            document.getElementById('f6Value').textContent = Data.data.Employee.Skill;
            document.getElementById('f7Value').textContent = Data.data.Employee.ID;
            document.getElementById('Bio').textContent = Data.data.Employee.Bio;
            document.querySelector('.bio').style.display = "flex";
            document.querySelector('.bio').style.flexDirection = "column";


            // Populate update form
            document.getElementById('updateName').value = Data.data.Name;
            document.getElementById('updateDOB').value = Data.data.Employee.BirthDate;
            document.getElementById('updateEmail').value = Data.data.Email;
            document.getElementById('updateLocation').value = Data.data.Employee.City;
            document.getElementById('updateF6').value = Data.data.Employee.Skill;
            document.getElementById('updateBio').value = Data.data.Employee.Bio;

        } else {
            //document.getElementById('heading').textContent = "EMPLOYER PROFILE";
            document.getElementById('name').textContent = Data.data.Name;
            document.getElementById('userId').textContent = Data.data.ID;
            document.getElementById('dob').textContent = Data.data.Employer.BirthDate;
            document.getElementById('email').textContent = Data.data.Email;
            document.getElementById('location').textContent = Data.data.Employer.City;
            document.getElementById('f6').textContent = "Company";
            document.getElementById('f7').textContent = "Employer ID";
            document.getElementById('f6Value').textContent = Data.data.Employer.Company;
            document.getElementById('f7Value').textContent = Data.data.Employer.ID;

            // Populate update form
            document.getElementById('updateName').value = Data.data.Name;
            document.getElementById('updateDOB').value = Data.data.Employer.BirthDate;
            document.getElementById('updateEmail').value = Data.data.Email;
            document.getElementById('updateLocation').value = Data.data.Employer.City;
            document.getElementById('updateF6').value = Data.data.Employer.Company;
            document.querySelector('.bio').style.display = 'none';
        }
    })
    .catch(error => {
        console.error('Error fetching data:', error);
    });
}

function showUpdateForm(role) {

    document.querySelector('.update-form').style.display = 'flex';
    document.querySelector('.update-form').style.flexDirection= 'column';


    document.querySelector('.content').style.display = 'none';

    if (role === "Employee") {
        document.querySelector('.bio').style.display = 'flex';
        document.getElementById('f6Label').textContent = "Skill";
    } else {
        document.getElementById('bio').style.display = 'none';
        document.getElementById('f6Label').textContent = "Company";
    }
}

function submitUpdateForm(id, token, role) {
    const updatedData = {
        Name: document.getElementById('updateName').value,
        BirthDate: document.getElementById('updateDOB').value,
        Email: document.getElementById('updateEmail').value,
        City: document.getElementById('updateLocation').value,
    };

    if (role === "Employee") {
        updatedData.Skill = document.getElementById('updateF6').value;
        updatedData.Bio = document.getElementById('updateBio').value;
    } else {
        updatedData.Company = document.getElementById('updateF6').value;
    }
    if (role === "Employee") {
    fetch(`http://127.0.0.1:8082/updateProfileEmployee/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(updatedData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Profile updated:', data);
        alert('Profile updated successfully!');
        window.location.reload();
    })
    .catch(error => {
        console.error('Error updating profile:', error);
    });
}
else{
    fetch(`http://127.0.0.1:8082/updateProfileEmployer/${id}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(updatedData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Profile updated:', data);
        alert('Profile updated successfully!');
        window.location.reload();
    })
    .catch(error => {
        console.error('Error updating profile:', error);
    });
}
}

function deleteProfile(id, token) {
    fetch(`http://127.0.0.1:8082/deleteUser/${id}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${token}`
        }
    })
    .then(response => {
        if (response.ok) {
            alert('Profile deleted successfully!');
            sessionStorage.removeItem('token');
            // Optionally, redirect to a different page or perform other actions
            window.location.href = 'index5.html';
        } else {
            throw new Error('Failed to delete profile');
        }
    })
    .catch(error => {
        console.error('Error deleting profile:', error);
    });
}
