<script>
  import { goto } from '$app/navigation'; // Import the SvelteKit navigation function
  import Comment from '../Comment.svelte'
  import CreatePostSvg from '../CreatePostSvg.svelte';
  import RightPanelFunctions from '../RightPanelFunctions.svelte';
  import Sidebar from '../Sidebar.svelte';
  import Notification from '../Notification.svelte'
  import Post from '../Post.svelte'
  import HomeSvg from '../HomeSvg.svelte';
  import LikeSvg from '../LikeSvg.svelte';
  import { onMount } from 'svelte'; // Import onMount for client-side operations


    let users = [];
    let firstNames = [];
    let lastNames = [];
    let albums = [];
    let photos = [];
    let showNotification = false;
    let showProfile = false;
    let selectedPostComments = []; 
    let selectedPostId = ""





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
    console.log("Value of the cookie: ", cookieValue);


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
  
//   export function DeleteCookie() {
//       var res = document.cookie;
//       var arrayCookie = res.split(";");
//       for(var i = 0; i < arrayCookie.length; i++) {
//           var key = arrayCookie[i].split("=");
//           document.cookie = key[0]+" =; expires = Thu, 01 Jan 1970 00:00:00 UTC";
//       }
//       goto('/')
//    }

   let posts = {
        AuthorId:cookieValue,
    }

   let comments = {
        AuthorId:cookieValue,
        // PostId: postId,
    }

   let user = {
        SessionId:cookieValue,
    }

    let resultPosts = []
    let resultComments = []
    export let resultUserInfo = []
    let sortedResultPosts = []

    async function AllFetches () {
        let result = await fetchCookieValue('http://localhost:8080/verificationSessionId', cookieValue)
        if (!result.hasOwnProperty("Success")) {
            goto('/');
            return
        } 
        console.log("Result :", result);

        let fetchedPosts = await fetchCookieValue('http://localhost:8080/getPost', posts);
        if (!fetchedPosts.hasOwnProperty("Success")) {
            goto('/');
            return;
        }

        let fetchedUser = await fetchCookieValue('http://localhost:8080/getUser', user);
        if (!fetchedUser.hasOwnProperty("Success")) {
            goto('/');
            return;
        }

        resultUserInfo = fetchedUser.userInfos[0]
        console.log("resltUserInfo: ",resultUserInfo);
        resultPosts = fetchedPosts.Posts
        console.log("Result Post . post: ", resultPosts);
        console.log("type of resultPosts: ", typeof(resultPosts))
        // console.log("Post dates before sorting:", resultPosts.map(post => post.CreationDate));
        // sortedResultPosts = resultPosts.sort(
        //     (a, b) => new Date(b.CreationDate).getTime() - new Date(a.CreationDate).getTime()
        // );
        sortedResultPosts = resultPosts.sort(
            (a, b) => new Date(b.CreationDate).getTime() - new Date(a.CreationDate).getTime()
        )
        console.log("sorted posts: ", sortedResultPosts);
        console.log("Result Comment: ", resultComments);
    
    }

    
      // Perform all fetches after the page is mounted in the browser
      onMount(async () => {
        // Retrieve the cookie value after the page has mounted
        cookieValue = getCookieValue();
        console.log("Value of the cookie: ", cookieValue);

        // Run all fetches only in the client (after page load)
        await AllFetches();
      });

    
    let COMMENTS = []

    const findUserName = (userId) => {
        const user = users.find(user => user.id === userId);
        return user ? user.name : 'Unknown User';
    };


    const findCommentPosts = (postId) => {
        return COMMENTS.filter(comment => comment.PostId === postId);
    }


    async function displaySectionComment(postId) {

        let fetchedComments = await fetchCookieValue('http://localhost:8080/getComment', comments);
        if (!fetchedComments.hasOwnProperty("Success")) {
            goto('/');
            return;
        }
        
        COMMENTS = fetchedComments.Posts
        console.log("COMMENTS: ", COMMENTS);
        selectedPostComments = findCommentPosts(postId);
        selectedPostId = postId
        showNotification = false
        let sectionComment = document.getElementById("commentSection")
        sectionComment.style.display = "flex"
    }

    // function hideSectionComment() {
    //     let sectionComment = document.getElementById("commentSection")
    //     sectionComment.style.display = "none"
    //     selectedPostComments = []
    // }

    // function bellIcon() {
    //     console.log("clickkkkk");
    //     selectedPostComments = []; 
    //     showNotification = !showNotification
    // }

    function handleToggleProfile(event) {
        const { showProfile: updatedShowProfile, showNotification: updatedShowNotification, selectedPostComments: updatedSelectedPostComments } = event.detail;
        showProfile = updatedShowProfile;
        showNotification = updatedShowNotification;
        selectedPostComments = updatedSelectedPostComments
    }

    let showModal = false;

    function toggleModal() {
        showModal = !showModal;
    }

</script>

<section class="flex h-screen bg-gray-200">
    <Sidebar { users } { albums } { photos } />
    <section class="flex h-screen w-full">
        <section class="flex flex-col w-[80%]">
            <section>
                <div class="bg-white flex justify-between items-center mt-4 m-2.5 p-0.5">
                    <div class="flex items-center align-middle">
                        <HomeSvg/>
                        <div class="flex items-center align-middle font-bold text-3xl text-blue-500 t-2 ml-2">Home</div>        
                    </div>
                    <div class="flex justify-center group items-center mr-2 cursor-pointer p-0.5 appearance-none" on:click={toggleModal}>
                        <input type="text" placeholder="Quoi de neuf ?" class="flex relative mr-2 w-32 cursor-pointer focus:outline-none focus:border-none" readonly> 
                        <CreatePostSvg/>
                    </div>
                        {#if showModal}
                      <Post {resultUserInfo} closeModal={toggleModal} />
                    {/if}
                </div>
                <!-- <textarea class="h-7 w-96 m-2" placeholder="Ecrire le contenu du post"></textarea> -->
                    <div class="bg-white m-2.5 h-[92vh] overflow-auto flex flex-col items-center
                              [&::-webkit-scrollbar]:w-2
                          [&::-webkit-scrollbar-track]:bg-gray-100
                          [&::-webkit-scrollbar-thumb]:bg-gray-300
                          [&::-webkit-scrollbar-track]:rounded-full
                          [&::-webkit-scrollbar-thumb]:rounded-full">
                        <!-- Render the fetched posts -->
                        {#if sortedResultPosts.length >= 0}
                            {#each sortedResultPosts as post}
                                <button on:click={displaySectionComment(post.Id)} id="postDiv" class="flex flex-col bg-white h-auto p-2 mt-4 mb-4 hover:scale-103 w-[95%]">
                                    <div class="flex flex-row justify-between p-4">
                                        {#if post.Username} 
                                            <div class="flex">Username: {(post.Username)}</div>
                                        {:else}
                                            <div class="flex">{(post.FirstName)} {(post.LastName)}</div>
                                        {/if}
                                        
                                        <div class="flex">
                                            <div class="flex ">
                                            <LikeSvg/>
                                            {post.LikeCount}</div>                                    
                                            <div class="w-4 h-5 bg-red-500">{post.DislikeCount}</div>
                                        </div>
                                        <!-- <div>Post ID: {post.id}</div> -->
                                    </div>
                                    <div class="flex p-2">Date: {post.CreationDate}</div>
                                    <div class="flex bg-gray-300 p-4">Body: {post.Text}</div>
                                    {#if post.Image}
                                        <div class="flex m-auto">
                                            <img src={post.Image} class="flex p-4 rounded-2xl w-[40vw] items-center" alt="">
                                        </div>
                                    {/if}
                                </button>
                            {/each}
                        {:else}
                            <p>Loading...</p>
                        {/if}
                    </div>
            </section>
        </section>
        <section id="commentSection" class="bg-white flex flex-col h-screen w-[35%] w-max-[40vw]">
            <RightPanelFunctions {showNotification} {selectedPostComments} on:toggleProfile={handleToggleProfile} />
            {#if !showNotification}
            <Comment {selectedPostComments} {selectedPostId} />
            {:else}
            <Notification />
            {/if}
    </section>
    </section>
</section>


