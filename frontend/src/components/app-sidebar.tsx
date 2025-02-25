import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
  } from "@/components/ui/sidebar"
import { Button } from "./ui/button";
import { GetAllFiles, OpenFileExplorer } from "../../wailsjs/go/main/App";
  
export function AppSidebar() {

  function openFileExplorer() {
    OpenFileExplorer() 
  }

  function getAllFiles() {
    const files = GetAllFiles()
    files.then(x => console.log(x))
  }

  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup />
        <SidebarGroup />
      </SidebarContent>
      <SidebarFooter>
        <Button onClick={openFileExplorer}>Open Folder</Button>
        <Button onClick={getAllFiles}>Get All Files</Button>
      </SidebarFooter>
    </Sidebar>
  )
}
  