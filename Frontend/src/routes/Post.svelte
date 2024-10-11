<script>
  import BellSvg from "./BellSvg.svelte";
  import ClosePostSvg from "./ClosePostSvg.svelte";
  import SparklesSvg from "./SparklesSvg.svelte";
  import { onMount } from 'svelte'; // Import onMount for client-side operations
  // import { cookieValue } from './home/+page.svelte';  // Adjust the import path to match your directory structure



    // Function to safely check if we are in the browser environment
    function isBrowser() {
        return typeof window !== 'undefined' && typeof document !== 'undefined';
    }

    // Function to get the cookie value, only run in the browser
    function getCookieValue() {
      if (isBrowser()) {
        console.log(document.cookie); // Safe access after the page is loaded
        let theCookies = document.cookie.split(";"); // Split cookies by semicolon
        for (let i = 0; i < theCookies.length; i++) {
          let cookie = theCookies[i].trim(); // Remove leading spaces
          console.log("cookie is:", cookie);
          if (cookie.startsWith("sessionId=")) {
            return cookie.substring("sessionId=".length); // Get the value after "sessionId="
          }
        }
      }
      return null; // Return null if sessionId cookie is not found
    }

    
    export let cookieValue = getCookieValue(); 

  console.log("Cookie Value in Post.svelte: ", cookieValue);


  export let resultUserInfo 
  export let closeModal;

  console.log("In Post component: ",resultUserInfo);

  let authorId = cookieValue
  let text = ""
  let image = ""
  // let postConfidentiality = ""
  let imageFile = "";  // Store the selected image file here
  let imageContent = "";  // Store the base64 content for preview and upload

  const date = new Date();

  let day = date.getDate();
  let month = date.getMonth() + 1;
  let year = date.getFullYear();
  // let creationDate = `${day}-${month}-${year}`;
  let creationDate = `${year}-${month}-${day}`;

  // console.log(creationDate); // "17-6-2022"


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
                document.getElementById('imagePreview').src = imageContent;
                document.getElementById('imagePreview').style.display = 'block'; // Show the image
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
        const imagePreveiw = document.getElementById("imagePreview")
        imagePreveiw.style.display = "flex"
        imagePreveiw.style.src = url


        // Clean up the object URL after use
        // imgElement.onload = function() {
        //     URL.revokeObjectURL(url);
        // };
    }

     function imageAvar() {
        const icon = document.getElementById("labelIcon")
        icon.click()
    }



    export async function fetchCookieValue(url, data) {
        console.log(url);
        console.log(data);
       try {
           const response = await fetch(url, {
               method: 'POST',
               body: JSON.stringify(data)
           });

           if (response.ok) {
               const result = await response.json();
               console.log("Result from fetch cookie value: ", result);
               console.log("url: ", url);
               return result
           } else {
               console.error('Cookie not valid');
           }
       } catch (error) {
           console.error('Error with the cookie:', error);
       }
   }

  async function createPost() {
    const PostData = {
      AuthorId: authorId,
        Text: text,
        Image: image,
        CreationDate: creationDate,
        Image: imageContent,
    }
    let fetchDataPostValue = fetchCookieValue("http://localhost:8080/createPost", PostData)
    console.log("Post Data from the post: ",PostData);
    console.log("Image du post: ", PostData.Image )
    if (!fetchDataPostValue.hasOwnProperty("Success")) {
      // goto('/')
      console.log("Error create post");
    } 
  }
    onMount(async () => {
        // Run all fetches only in the client (after page load)
        await createPost();
     });

  let publishButton = false

  // reactive statement that updates publishButton based on whether the textarea has any non-whitespace text
  $: publishButton = text.trim() !== "";




</script>

<div class="fixed z-10 inset-0 bg-gray-600 bg-opacity-75 flex justify-center items-center">
  <div class="bg-white p-6 rounded-lg shadow-lg w-[50vw] h-auto">
  <div class="flex items-center text-center align-middle mb-4 justify-around">
    <div></div>
    <div class="flex justify-center">
        <h2 class="text-3xl font-bold">Créer une publication !</h2>
        <SparklesSvg/>
    </div>
    <button on:click={closeModal} class="flex cursor-pointer">
        <ClosePostSvg/>
    </button>
  </div>
  <div class="flex h-[10vh] items-center">
    <div class=" flex justify-center w-[8vw]">
        <img class="h-[60px] w-[60px] rounded-lg bg-slate-200" src="./basic-avatar.png" alt="Image for the group">
    </div>
    <div class=" w-full">
        <div class="flex p-0.5">{resultUserInfo.FirstName} {resultUserInfo.LastName}</div>
        <div class="flex p-0.5"> Amis
           <!-- <select
              class="rounded-md p-1 border-2 border-gray-500"
              name="ChooseConfidentiality"
              id="ChooseConfidentiality"
              bind:value={postConfidentiality}
              >
              <option value="">Sélectionnez la confidentialité</option>
              <option value="Public">Publique</option>
              <option value="Private">Privé</option>
              <option value="Almost privte">Presque privé</option>
          </select> -->
        </div>
    </div>
  </div>
  <div class="flex flex-col items-center justify-center">
    <textarea id="postText" class="bg-gray-100 flex justify-center items-center resize-none w-[45vw] h-[20vh]" bind:value={text}></textarea>
    <!-- <div class="bg-blue-700" id="imageDiv">
          <input
            type="file"
            accept="image/*"
            id="userFile"
            name="image"
            class="p-1.5 text-sm flex"
            on:change={handleFileChange}
          />
    </div> -->
        <div id="imageDiv" class="flex flex-col items-center m-3 relative">
            <!-- <img id="image" src="" alt="" on:click={imageAvar} class="flex justify-center w-32 h-32 object-cover rounded-full cursor-pointer" /> -->

            <input
                type="file"
                accept="image/*"
                id="userFile"
                name="image"
                class="p-1.5 text-sm flex"
                on:change={handleFileChange}
            />
        </div>
        <img id="imagePreview" class="hidden" alt="">
  </div>
    {#if publishButton}
      <button id="publishButton" on:click={createPost} on:click={closeModal}
        class="mt-10 bg-blue-500 text-white px-4 py-2 text-lg rounded w-full">
        Publier
      </button>
    {/if}
  </div>
</div>

PRIVATE PUBLIC ALMOST PRIVATE