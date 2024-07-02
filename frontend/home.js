document.addEventListener('DOMContentLoaded', function() {
    const token=sessionStorage.getItem('token');
    console.log(token);
    if(token){
        const decoded=jwt_decode(token);
        console.log(decoded);
        if(decoded.role=="Employer"){
            const postButtonHolder=document.getElementById('postButton-holder');
            const postButton=document.createElement('button');
            postButton.id="postBt"
            postButton.type="submit";
            postButton.textContent="Post Job";
            postButtonHolder.appendChild(postButton);
            postButton.style.paddingLeft="25px";
            postButton.style.paddingRight="25px";
            postButton.style.paddingTop="8px";
            postButton.style.paddingBottom="8px";
            postButton.style.fontSize="1.2rem";
            postButton.style.backgroundColor="#407ff0"
            postButton.style.color="whitesmoke";
            postButton.style.borderRadius="0.5vw"
            postButton.style.margin="2vw"
            document.getElementById('posted').textContent="Jobs Posted";

            document.getElementById('postBt').addEventListener('click',function(event) {
                event.preventDefault();
                window.location.href='http://127.0.0.1:3002/frontend/postJob.html';
            })

    }
}
else{
    alert("Please login");
    window.location.href='http://127.0.0.1:3002/frontend/index5.html';
}
});

            
        
    


    //}
      





