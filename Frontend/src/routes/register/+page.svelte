<script>
    import { goto } from '$app/navigation'; // Import the SvelteKit navigation function
    import Header from '../Header.svelte';


    // Creating variables to store data 
    let Password = '';
    let ConfirmPassword = '';
    let Username = '';
    let Firstname = '';
    let Lastname = '';
    let BirthDate = '';
    let Mail = '';
    let SelectGender = '';
    let imageSrc = ""
    let imageFile;
    let imageContent = ''; // Variable to hold the file content


    function handleFileChange(event) {
        const [file] = event.target.files;
        if (file) {
            imageFile = file;
            console.log('File Name:', file.name); 
            console.log('File Type:', file.type);  // Logs the file type
            console.log('File Size:', file.size);  // Logs the file size in bytes

            // Use FileReader to read the content of the file
            const reader = new FileReader();
            reader.onload = () => {
                imageContent = reader.result; // This contains the base64 encoded content or raw content
                console.log('File Content:', imageContent); // You can also use this to preview the image
                console.log(imageContent.length); // You can also use this to preview the image
                decodeImage(imageContent, document.getElementById("image")) // The second argument necessary refers to an image tag !!
            };
            reader.readAsDataURL(file); // Reads the file as a data URL
        }
    }


    function decodeImage(base64Image, imageDiv) {

        // Step 1: Extract the Base64 data
        base64Image = base64Image.split(',')[1]; // Get only the Base64 part

        // Step 2: Decode the Base64 string
        const binaryString = atob(base64Image);

        // Step 3: Convert the binary string to a Uint8Array
        const len = binaryString.length;
        const bytes = new Uint8Array(len);
        for (let i = 0; i < len; i++) {
            bytes[i] = binaryString.charCodeAt(i);
        }

        // Optional: Create a Blob object for the image
        const blob = new Blob([bytes], { type: 'image/jpeg' });

        // Display the image in an <img> element
        const url = URL.createObjectURL(blob);
        imageDiv.src = url;

        // Clean up the object URL after use
        imgElement.onload = function() {
            URL.revokeObjectURL(url);
        };
    }

    function imageAvar() {
        const icon = document.getElementById("labelIcon")
        icon.click()
    }

 
    /**
     * Function to send data to the back end server in go
     *
     * Fetch the data and wait the response after the check in go
     * @param event
     */
    async function handleSubmit(event) {
        event.preventDefault(); // Prevent the default form submission
 
        // Check if both password and ConfirmPassword are the same
        if (Password !== ConfirmPassword) {
            console.error('Passwords do not match');
            return;
        }
 
        // Const to store the data from the user
        const formData = {
            Username: Username,
            FirstName: Firstname,
            LastName: Lastname,
            BirthDate: BirthDate,
            Email: Mail,
            Gender: SelectGender,
            Password: Password,
            ConfirmPassword: ConfirmPassword,
            ProfilePicture: imageSrc
        };
        console.log(formData);
        console.log(imageSrc);

 
        try {
            const response = await fetch('http://localhost:8080/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            });
 
            if (response.ok) {
                const result = await response.json();
                console.log('Registration successful:', result);
                if (result.hasOwnProperty("Success")) {
                // Print in the console the data if the registration has been a success
                console.log('File and data uploaded successfully:', result);
                //Reset of the variables
                 Password = '';
                 ConfirmPassword = '';
                 Username = '';
                 Firstname = '';
                 Lastname = '';
                 BirthDate = '';
                 Mail = '';
                 SelectGender = '';
                 ProfilePicture = '';
                goto('/');
               } else {
                return
               }
            } else {
                console.error('Failed to register');
            }
        } catch (error) {
            console.error('Error during registration:', error);
        }

}
 </script>
    <!--Call Header component  -->
	<Header />

 <div class="flex justify-center">
 <form class="flex flex-col p-5 w-auto md:w-auto sm:w-auto lg:w-[25vw]" on:submit={handleSubmit} enctype="multipart/form-data">
     <div class="flex m-3 text-blue-500 text-3xl font-bold text-left">Create an account</div>
 
     <input
         type="text"
         id="Username"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Username"
         bind:value={Username}
     />
     <input
         type="text"
         id="Firstname"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Firstname"
         bind:value={Firstname}
     />
     <input
         type="text"
         id="Lastname"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Lastname"
         bind:value={Lastname}
     />
     <input
         type="date"
         id="BirthDate"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         max="2024-12-31"
         bind:value={BirthDate}
     />
     <select
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         name="SelectGender"
         id="SelectGender"
         bind:value={SelectGender}
     >
         <option value="">Select Gender</option>
         <option value="Homme">Homme</option>
         <option value="Femme">Femme</option>
         <option value="Ne pas répondre">Ne pas répondre</option>
     </select>
     <input
         type="email"
         id="Mail"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Email"
         bind:value={Mail}
     />
     <input
         type="password"
         id="Password"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Password"
         bind:value={Password}
     />
     <input
         type="password"
         id="ConfirmPassword"
         class="m-3 rounded-md p-1.5 border-2 border-gray-500"
         placeholder="Confirm Password"
         bind:value={ConfirmPassword}
     />

    <div id="imageDiv" class="flex flex-col items-center m-3 relative">
            <img id="image" src="./basic-avatar.png" alt="avatar pic" on:click={imageAvar} class="flex justify-center w-32 h-32 object-cover rounded-full cursor-pointer" />

            <input
                type="file"
                accept="image/*"
                id="userFile"
                name="image"
                class="p-1.5 text-sm hidden"
                on:change={handleFileChange}
            />
            <label id="labelIcon" for="userFile" class="absolute -top-4  text-white bg-cyan-500 w-7 h-7 flex items-center justify-center rounded-full text-lg cursor-pointer ">
              <span class="material-symbols-outlined"> photo_camera </span>
            </label>
        </div>
 
     <button type="submit" class="ml-auto m-3 p-1 border-2 rounded-md bg-blue-500 text-white w-40">
         Sign up baby
     </button>
 </form>
 </div>

 