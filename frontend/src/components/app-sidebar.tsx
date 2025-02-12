import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
  } from "@/components/ui/sidebar"
import { Button } from "./ui/button";
import { OpenFileExplorer } from "../../wailsjs/go/main/App";
  
  export function AppSidebar() {

    function openFileExplorer() {
      OpenFileExplorer()  
    }

    return (
      <Sidebar>
        <SidebarContent>
          <SidebarGroup />
          <SidebarGroup />
        </SidebarContent>
        <SidebarFooter>
          <Button onClick={openFileExplorer}>Open Folder</Button>
        </SidebarFooter>
      </Sidebar>
    )
  }
  