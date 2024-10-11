<script>
    import { onMount } from 'svelte'; // Import onMount for client-side operations

    export let selectedPostComments = []; 
    export let selectedPostId = "";
    let showCommentBox = false;

    function toggleCommentBox() {
        showCommentBox = !showCommentBox;
        console.log("Comment box toggle!", showCommentBox);
    }


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

    console.log("In Comment component: ",resultUserInfo);

    let authorId = cookieValue
    export let text = ""    
    const date = new Date();
    let day = date.getDate();
    let month = date.getMonth() + 1;
    let year = date.getFullYear();
    // let creationDate = `${day}-${month}-${year}`;
    let creationDate = `${year}-${month}-${day}`;

    // console.log(creationDate); // "17-6-2022"



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
    
    export async function CreateComment() {
        const CommentData = {
            AuthorId: authorId,
            Text: text,
            CreationDate: creationDate,
            PostId: selectedPostId,
        }

        let fetchDataPostValue = fetchCookieValue("http://localhost:8080/createComment", CommentData)
        console.log("Post Data from the post: ",CommentData);
        console.log("Image du post: ", CommentData.Image )
        if (!fetchDataPostValue.hasOwnProperty("Success")) {
          // goto('/')

          console.log("Error create post");
        } 
     }


     onMount(async () => {
        // Run all fetches only in the client (after page load)
        await CreateComment();
      });

      let publishButton = false

      // reactive statement that updates publishButton based on whether the textarea has any non-whitespace text
      $: publishButton = text.trim() !== "";


    //   console.log("Selected comments: ", selectedPostId);

</script>

    <div class="h-screen overflow-hidden relative">
        <!-- <button on:click={hideSectionComment}>X</button> -->
        <div class="flex flex-col overflow-auto h-[70vh]">
            <!-- Render the comments for the selected post -->
            {#if selectedPostComments.length > 0}
                    <span class="font-bold ml-2">Commentaires</span>

                {#each selectedPostComments as comment}
                    <div class="flex flex-col bg-gray-300 h-auto p-1 m-3 text-xs text-center">
                        <div class="flex flex-col justify-around p-4 w-[90%]">  
                            <!-- <div><strong>Name:</strong> {comment.name}</div>
                            <div><strong>Email:</strong> {comment.email}</div> -->
                            <div><strong>Author Id: </strong> {comment.AuthorId}</div>
                            <div><strong>Text</strong> {comment.Text}</div>
                            <div><strong>Creation Date: </strong> {comment.CreationDate}</div>
                            <div><strong>Like: </strong> {comment.LikeCount}</div>
                            <div><strong>Dislike: </strong> {comment.DislikeCount}</div>
                            <div><strong>Dislike: </strong> {comment.PostId}</div>
                        </div>
                    </div>
                {/each}
            {:else}
                <p class="flex justify-center text-center text-xl">No  comments available, please select a post.</p>
            {/if}
             <!-- Comment Button -->
        </div>
        <div class="flex w-auto justify-center">
            <button on:click={toggleCommentBox} class="bg-blue-500 flex justify-center w-[10vw] rounded-lg text-white mt-3 mb-3">
                Commenter
            </button>
        </div>
        <!-- Comment Box (appears from the bottom, inside the section) -->
        <div class="flex flex-col absolute bottom-0 p-4 transition-transform duration-500 bg-gray-100  h-[15vh] w-full overflow-auto"
            style="transform: translateY({showCommentBox ? '0' : '100%'})">
            <textarea bind:value={text} class="w-full p-2 border border-gray-300 rounded-md resize-none" placeholder="Ecrire un commentaire..." ></textarea>
            {#if publishButton}
                <div class="flex justify-end mt-2">
                    <button  on:click={CreateComment}  on:click={toggleCommentBox} on class="bg-blue-500 text-white rounded-lg p-1">Envoyer</button>
                 </div>
            {/if}
        </div>
    </div>
