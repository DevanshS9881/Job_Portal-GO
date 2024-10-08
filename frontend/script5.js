function checkLoginStatus() {
    if (localStorage.getItem('token') != null) {
        window.location.href = 'homepage.html';
    }
}
window.onload = function() {
    // Check if the user is already logged in
     if (localStorage.getItem('token') != null) {
        alert("Session is already logged in");
        window.location.href = 'homepage.html';
    }
    
    // Disable caching of the login page to ensure the back button doesn't show the cached page
    window.addEventListener('pageshow', function(event) {
        if (event.persisted) {
            window.location.reload(); 
        }
    });
    window.addEventListener('popstate', checkLoginStatus);
    window.addEventListener('hashchange', checkLoginStatus);
};
const loginBox=document.querySelector('.loginBox');
const loginAsk=document.querySelector('.loginAsk');
const signupAsk=document.querySelector('.signupAsk');
loginAsk.addEventListener('click',function(e){
    e.preventDefault();
    loginBox.classList.add('active');
})
document.querySelector('.signupAsk').addEventListener('click', function(e) {
    e.preventDefault();
    document.querySelector('.loginBox').classList.remove('active');
});

document.getElementById('signupBt').addEventListener('click', async function(event) {
    event.preventDefault();
    const name=document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    const response = await fetch('https://code-backend-backend.onrender.com/register', { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name,email, password })
    });
    console.log(response)
    if(response.ok){
        console.log("Fetched the api successfully");
    }

    const data = await response.json();
    if (response.ok) {
        localStorage.setItem('token', data.token); 
        alert('Signup and login successful!');
        window.location.href = 'protected.html'; 
    } else {
        alert('Signup failed: ' + data.message);
    }
});

// document.getElementById('googleUp').addEventListener('click', function() {
//       window.location.href = 'https://code-backend-backend.onrender.com/google_login'; // Update this URL to match your backend's Google OAuth URL
//  });
document.getElementById('loginBt').addEventListener('click', async function(event) {
    event.preventDefault()
    const email = document.getElementById('emailLogin').value;
    const password = document.getElementById('passwordLogin').value;

    const response = await fetch('https://code-backend-backend.onrender.com/login', { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    });
   
    const data = await response.json();
    if (response.ok) {
        localStorage.setItem('token', data.token); 
        alert('Signup and login successful!');
        window.location.href = 'homepage.html'; 
    } else {
        alert('Signup failed: ' + data.message);
    }
});
document.getElementById('googleUp').addEventListener('click', async function(event) {
    event.preventDefault();
    window.location.href = 'https://code-backend-backend.onrender.com/google_login'; 
    const params = new URLSearchParams(window.location.search);
    const code = params.get('code');
    const state = params.get('state');
                try {
                    const response = await fetch('https://code-backend-backend.onrender.com/google_callback?code=' + code + '&state=' + state,{
                        method: 'GET',
                        credentials: 'include' 
                    });
                
                            //     method: 'GET',
                            //     headers: {
                            //         'Content-Type': 'application/json'
                            //     },
                            // });
                    const data = await response.json();
                    console.log(data.token);

                    if (data.token) {
                        
                        localStorage.setItem('token', data.token);

                        //Redirect to protected page
                       window.location.href = 'protected.html';
                    } else {
                        console.error('Token not found in response:', data);
                    }
                } catch (error) {
                    console.error('Failed to fetch token:', error);
                }
            // }
        });







//     const response = await fetch('https://code-backend-backend.onrender.com/google_callback', { 
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json'
//         },
//     });
//     const data = await response.json();
//     if (response.ok) {
//         localStorage.setItem('token', data.token); 
//         alert('Signup and login successful!');
//         window.location.href = '/protected.html'; 
//     } else {
//         alert('Signup failed: ' + data.message);
//     }
// });
// ;