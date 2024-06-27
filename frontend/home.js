document.addEventListener('DOMContentLoaded', function() {
    // Replace the URL with your actual endpoint
    // const token=sessionStorage.getItem('token');
    // console.log(token);
    // //if(token){
    //     const decoded=jwt_decode(token);
    //     console.log(decoded);
        // if(decoded.Role=="Employer"){
            const postButtonHolder=document.getElementById('postButton-holder');
            const postButton=document.createElement('button');
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


        })


            
        
    


    //}
      





