<script lang="ts">
    let username: string | undefined = $state();
    let password: string | undefined = $state();

    function login(): void {
        if (username === undefined) {
            // TODO: Client-side error
            return;
        }

        if (password === undefined) {
            // TODO: Client-side error
            return;
        }

        fetch("/api/login", {
            method: "GET",
            redirect: "manual",
            headers: {
                Authorization: "Basic " + btoa(username + ":" + password)
            }
        })
        .then((response) => {
            if (response.status === 200) {
                window.location.href = "/";
                return;
            }
            console.log(response);
        })
        .catch((error: Response) => console.log(error));
    }

    function register(): void {
        fetch("/api/register", {
            method: "POST",
            body: JSON.stringify({username, password})
        })
        .then((response: Response) => {
            if (response.status === 201) {
                login();
                return;
            }
            console.log(response);
        })
        .catch((error) => console.log(error));
    }
</script>

<div class="flex h-[100vh] w-[100vw] justify-center flex-col">
    <fieldset class="fieldset bg-base-200 rounded-box w-md flex flex-col ml-auto mr-auto p-10 gap-5">
        <h1 class="text-3xl text-center -mt-5">Lorem ipsum</h1>
        <label class="input w-[100%]">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z" />
            </svg>
            
            <input id="username" class="grow" placeholder="Usuário" bind:value={username} />
        </label>
        <label class="input w-[100%]">
            
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z" />
            </svg>
            <input id="password" class="grow" placeholder="Senha" type="password" bind:value={password} />
        </label>
        
        <div class="w-full flex justify-between">
            <button class="btn btn-primary w-[49%]" onclick="{login}">Entrar</button>
            <button class="btn btn-secondary w-[49%]" onclick="{register}">Registrar</button>
        </div>
    </fieldset>
</div>