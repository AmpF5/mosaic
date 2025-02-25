// src/components/MainContent.tsx
import React, { useEffect, useState } from "react";
import { Button } from "./ui/button";
import { GetAllFiles } from "../../wailsjs/go/main/App";
import {database} from '../../wailsjs/go/models';



const MainContent: React.FC = () => {
  const [files, setFiles] = useState<database.File[]>([]);

  useEffect(() => {
    GetAllFiles()
      .then(x => setFiles(x))
      .catch(error => console.error("Error fetching files",error))
  }, [])
  
  return (
    <div className="grid auto-cols-max grid-cols-3 gap-4">
      {files.map((file, index) => (
        <div key={index} className="p-4 bg-gray-200 rounded-lg">
          <img src={`file://${file.FilePath.replace(/\\/g, "/")}`} alt="Preview" className="w-32 h-32 object-cover rounded-lg" />
          {file.FilePath}
        </div>
      ))}
    </div>
  );
};

export default MainContent;
