export const serverMainScript = `import { world } from "@minecraft/server";

world.beforeEvents.worldInitialize.subscribe(() => {
    world.sendMessage("Hello world!");
});`;