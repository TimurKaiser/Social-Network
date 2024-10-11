<script>
  import Logo from './Logo.svelte';
  import { goto } from '$app/navigation'; // Import the SvelteKit navigation function


    // export {deleteCookie} from './home/+page.svelte'
    export let users = [];
    export let albums = [];
    export let photos= [];
</script>

<section class="relative flex h-screen">
    <section class="flex flex-col bg-gray-800 h-sreen w-20 ">
    <button on:click={()=> goto("/home")} class="ml-2 z-0 cursor-pointer"> 
        <Logo/>
    </button>
        {#each users as user, userIndex (userIndex)}
        {#each albums as album, albumIndex (albumIndex)}
            {#if album.userId == user.id && (albumIndex == 0 || album.userId !== albums[albumIndex - 1].userId)}
            {#each photos as photo, photoIndex (photoIndex)}
            {#if photo.albumId == album.id && (photoIndex == 0 || photo.albumId !== photos[photoIndex - 1].albumId)}
              <div class="group relative flex items-center justify-center">
                <img class="rounded-full max-w-[64px] m-1" src={photo.url} alt="Album photo" />
                <!-- User name appears on hover -->
                <p class="absolute left-full bg-gray-900 text-white text-sm rounded px-7 py-1 opacity-0 group-hover:opacity-100 transition-opacity">
                    {user.name}
                </p>
            </div>
            {/if}
            {/each}
            {/if}
        {/each}
        {/each}
        <a href="/message" class="rounded-full bg-blue-500 text-white max-w-[64px] min-h-[64px] m-1 mt-10 text-center items-center justify-center flex"><button>+</button></a>
        <a href="/profile" class="rounded-full bg-yellow-600 text-white max-w-[64px] min-h-[64px] m-1 mt-10 text-center items-center justify-center flex"><button>Profile</button></a>
        <!-- <button class="rounded-full bg-red-600 m-1 text-sm p-1" on:click={deleteCookie}>Log out</button> -->
    </section>
    <section class="flex flex-col bg-zinc-00 h-screen items-center w-[90%]">
        <!-- <input class="ml-2 text-2xl  text-blue-500 my-4 ms-20 max-w-[px] p-2" placeholder="Search..."> -->
        <div class="flex h-[7vh]">
            <input class="text-xxl  text-blue-500 m-4 max-w-[px] p-1" placeholder="Search...">
        </div>
        <div class="bg-white flex flex-col h-[30vh] w-[90%] justify-center items-center mt-4 mb-4">Evénements</div>
        <div class="flex flex-col bg-white mt-4 mb-4 h-[50vh] w-[90%] justify-center items-center">
            <div class="flex flex-col">
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="flex margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>
                </div>
                <div class="flex flex-row m-2">
                    <span class="h-[25px] w-[25px] bg-customPP rounded-full mr-4"></span>
                    <p class="margin-inline-start">Nom de Groupe</p>        
                </div>
            </div>
        </div>
        <div>
            <a href="/group" class="bg-blue-500 p-2 rounded-xl text-white"><button class="w-36 mt-6">New group</button></a>
        </div>
    </section>
</section>