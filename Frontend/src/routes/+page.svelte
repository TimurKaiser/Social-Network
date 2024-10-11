<script> 
   import { goto } from '$app/navigation'; // Import the SvelteKit navigation function
   import Header from './Header.svelte';


    // Creating variables to get the data from the user 
   let valueMail;
   let valuePassword
   let disabled = true



    async function handleLogin() {
        const formDataLogin = {
        Email: valueMail,
        Password: valuePassword
        }

       try {
           const response = await fetch('http://localhost:8080/login', {
               method: 'POST',
               body: JSON.stringify(formDataLogin)
           });

           if (response.ok) {
               const result = await response.json();
                if (result.hasOwnProperty("Success")) {
                console.log('File and data uploaded successfully:', result);
                //Setting the cookie if the user is already registered
                document.cookie = `sessionId=${result.sessionId}; path=/`;
                valueMail = "";
                valuePassword = "";
                goto('/home');
               } else {
                console.log("Error");
                const errorDiv = document.getElementById("divError")
                errorDiv.style.display = "flex"
                return
               }
           } else {
               console.error('Failed to upload file');
           }
       } catch (error) {
           console.error('Error uploading file:', error);
       }
   }

   

</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />

    
</svelte:head>

	<Header />


<section>
    <form class="flex flex-col p-5 w-auto md:w-auto sm:w-auto lg-[25vw]" on:submit={handleLogin}>
        <div class="flex m-3 justify-center text-blue-500 text-3xl font-bold text-center">Login with an account</div>
        <div class="flex flex-col m-4 justify-center text-black">
            <input
                type="text"
                id="Username"
                class="m-3 p-1.5 rounded-md border-2 border-gray-500"
                placeholder="Mail"
                required
                bind:value={valueMail}
            />

            <input
                type="password"
                id="Password"
                class="m-3 p-1.5 rounded-md border-2 border-gray-500"
                required
                placeholder="********"
                bind:value={valuePassword}
            />
            <button
                id="LoginButton"
                class="m-3 border-2 p-1 rounded-md bg-blue-500 text-white"
                type="submit"
                disabled={valueMail === "" || valuePassword === ""}
            >
                Login
            </button>
            <a
                href="/register"
                class="m-3 p-1 justify-center flex no-underline hover:no-underline border-2 rounded-md bg-blue-500 text-white"
            >
                <button id="ToRegisterButton">Register</button>
            </a>

            <a
                href="/home"
                class="m-3 p-1 justify-center flex no-underline hover:no-underline border-2 rounded-md bg-blue-500 text-white"
            >
                <button id="ToRegisterButton">Access to blog</button>
            </a>

            <div id="divError" class="m-3 p-1 justify-center hidden no-underline hover:no-underline border-2 rounded-md bg-red-600 text-white">
                <span>Mauvais mot de passe ou mauvaise adresse mail</span>
            </div>

        </div>
    </form>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
</style>
