<script>
    import About from './About.svelte';

    export let users = [];
    let isEditing = false;
    
    let isOpen = false;
    let isNotif = false;

    function toggleEditMode() {
        isEditing = !isEditing;
    }

//     async function getFollower() {

//        try {
//            const response = await fetch('http://localhost:8080/getFollower', {
//                method: 'POST',
//                body: JSON.stringify(formDataLogin)
//            });

//            if (response.ok) {
//                const result = await response.json();
//                 if (result.hasOwnProperty("Success")) {
//                 console.log('File and data uploaded successfully:', result);
//                 //Setting the cookie if the user is already registered
//                 document.cookie = `sessionId=${result.sessionId}; path=/`;
//                 valueMail = "";
//                 valuePassword = "";
//                 goto('/home');
//                } else {
//                 return
//                }
//            } else {
//                console.error('Failed to upload file');
//            }
//        } catch (error) {
//            console.error('Error uploading file:', error);
//        }
//    }

</script>

        <section class="flex flex-col h-screen">
            <div class="flex flex-col h-[15vh] mb-8 relative">
                <div class="bg-[#0094FF] h-[60%] flex relative"></div>
                <div class="flex group absolute top-8 left-10">
                    <img class="h-[120px] w-[120px] rounded-lg bg-slate-400">
                    <span class="flex text-sm ml-2 h-[3vh] items-center font-bold relative top-16">
                        {#each users as user}
                            {#if user.id == 1}
                                {user.name}
                            {/if}
                        {/each}
                    </span>
                </div>
            </div>
            <div>
                <div>
                    <span class="flex font-bold ml-2">About</span>
                    <div class="flex justify-center">
                        <About {toggleEditMode} {users} isEditing={isEditing} />
                    </div>
                </div>
            </div>
            <div class="flex flex-row justify-center">
                <!-- <button on:click{getFollower} class="m-2 w-[86px] h-[32px] bg-[#D9D9D9] text-black font-inter text-base font-bold leading-normal">Abonnés</button> -->
                <button class="m-2 w-[124px] h-[32px] bg-[#D9D9D9] text-black font-inter text-base font-bold leading-normal">Abonnés</button>
                <button class="m-2 w-[124px] h-[32px] bg-[#D9D9D9] text-black font-inter text-base font-bold leading-normal">Abonnements</button>
            </div>

            <!-- Limit the height of the members list and allow it to scroll -->
            <div class="flex flex-col mt-4 mb-4 w-full overflow-y-auto h-[400px]">
                <div class="flex flex-col">
                    {#each users as user}
                        <div class="flex flex-row m-2">
                            <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                            <p class="bg-customGray flex items-center w-full p-2">Membre {user.id}</p>
                        </div>
                    {/each}
                </div>
            </div>
        </section>
