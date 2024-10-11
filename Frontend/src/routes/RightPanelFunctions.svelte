<script>
    import { createEventDispatcher } from 'svelte'; // Create a dispatcher
    import BellSvg from './BellSvg.svelte';
    import LogoutSvg from './LogoutSvg.svelte';
    import ProfileSvg from './ProfileSvg.svelte';
     import { goto } from '$app/navigation'; // Import the SvelteKit navigation function


    export let selectedPostComments = [];
    export let showNotification = false;
    export let showProfile = true;

    const dispatch = createEventDispatcher(); // Initialize dispatcher

    let isOpen = false;

    function toggleDropdown() {
        isOpen = !isOpen;
    }

    function bellIcon() {
        selectedPostComments = [];
        let toggleProfile
        let toggleNotification
        if (!showProfile && !showNotification){
            toggleProfile = false
            toggleNotification = true
        } else {
            toggleProfile = !showProfile;
            toggleNotification = !showNotification && !toggleProfile;
        }
        dispatch('toggleProfile', { showProfile: toggleProfile, showNotification: toggleNotification, selectedPostComments: selectedPostComments });
    }

      function DeleteCookie() {
      var res = document.cookie;
      var arrayCookie = res.split(";");
      for(var i = 0; i < arrayCookie.length; i++) {
          var key = arrayCookie[i].split("=");
          document.cookie = key[0]+" =; expires = Thu, 01 Jan 1970 00:00:00 UTC";
      }
      goto('/')
   }




    
</script>

<div class="flex flex-row m-2.5">
    <input class="bg-customGray text-xxl text-black m-4 max-w-[px] p-1" placeholder="Search...">
    <div class="flex items-center align-middle">
        <ul class="flex flex-row items-center h-[7vh] justify-center">
            <li class="mr-2 cursor-pointer">
                <button on:click={toggleDropdown} class="bg-yellow-400 rounded-full p-1 w-[32px] m-1">O</button>
                {#if isOpen}
                <ul class="absolute z-10 right-0 mt-2 w-48 bg-[#6C6C6C] rounded shadow-lg">
                    <li class="px-4 py-2 hover:bg-gray-200">
                        <div class="flex items-center">
                            <ProfileSvg/>
                            <a href="/profile" class="text-white">Profile</a>
                        </div>
                    </li>
                    <li class="px-4 py-2 hover:bg-gray-200">
                        <div class="flex items-center">
                            <LogoutSvg/>
                            <button on:click={DeleteCookie} class="text-white">Logout</button>
                        </div>
                    </li>
                </ul>
                {/if}
            </li>
            <button on:click={bellIcon} class="ml-2 z-10 cursor-pointer">
                <BellSvg/>
            </button>
        </ul>
    </div>
</div>
