"use client"

import { useEffect, useState } from "react"
import VideoCard from "../components/VideoCard"
import Pagination from "../components/Pagination"

interface Video {
  id: string;
  title: string;
  description: string;
  published_at: string;
  thumbnail: string;
  url: string
}

interface VideoPageResponse {
  videos: Video[];
  currentPage: number;
  totalPages: number;
  sortedIn: string;
  filter: Filter;
}

interface Filter {
  title: string;
  start: number;
  end: number;
}

export default function Home() {
  const [currentPage, setCurrentPage] = useState(1)
  const [currentVideos, setCurrentVideos] = useState<VideoPageResponse>({
    videos: [],
    currentPage: 1,
    totalPages: 1,
    sortedIn: "published_at_desc",
    filter: {
      title: "",
      start: 0,
      end: Date.now(),
    }
  })
  const [totalPages, setTotalPages] = useState<number>(0)
  const [currSort, setCurrSort] = useState<string>("published_at_desc")
  const [inpValue, setInpValue] = useState<string>("")
  const [filter, setFilter] = useState<Filter>({
    title: "",
    start: 0,
    end: Date.now()
  })

  useEffect(() => {
    fetch(`/api/videos?page=${currentPage}&sort=${currSort}&title=${filter?.title}&start=${filter?.start}&end=${filter?.end}`).then((data) => {
      if (data.status != 200) {
        return data.text().then((text) => {
          alert(`Error: ${text}`);
        });
      }
      return data.json()
    }).then((data: VideoPageResponse) => {
      setCurrentVideos(data)
      setTotalPages(data.totalPages)
    })
    
  }, [currentPage, currSort, filter])

  const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setCurrSort(event.target.value);
  };

  const updateFilter = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInpValue(event.target.value)
  };

  const applyFilter = () => {
    var result = Object.fromEntries(
        inpValue.split('&').map(p => p.split('='))
    )

    var filter: Filter = {
      title: result.title === undefined ? "" : result.title,
      start: result.start === undefined ? 0 : Date.parse(result.start),
      end: result.end === undefined ? Date.now() : Date.parse(result.end)
    }

    if (isNaN(filter.start) || isNaN(filter.end)) {
      alert("Invalid start or end time, please recheck")
    } else {
      setFilter(filter)
    }

    var inp = document.getElementById("filter") as HTMLInputElement;
    inp.value = ""
    setInpValue("")
  }

  return (
    <main className="min-h-screen bg-background py-12 px-4 sm:px-6 lg:px-8">
      <h1 className="text-4xl font-bold text-center text-primary mb-12">Latest Cats Videos</h1>
      <div>
        <label htmlFor="sort-opt">Sort by </label>
        <select id="sort-opt" name="options" value={currSort} onChange={handleChange} className="p-2 mb-7">
          <option value="published_at_desc">Publish time</option>
          <option value="title_asc">Title</option>
        </select>
        <input type="text" name="filter" id="filter" className="ml-5 mr-5 p-2 w-80" placeholder="title=example&start=mm-dd-yyyy" onChange={updateFilter} />
        <button className="bg-cyan-200 rounded-[0.2vw] font-semibold p-2 text-slate-700" onClick={applyFilter}>Apply</button>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        {currentVideos?.videos.map((video) => (
          <VideoCard key={video.id} title={video.title} thumbnailUrl={video.thumbnail} videoUrl={video.url} />
        ))}
      </div>
      <Pagination currentPage={currentPage} totalPages={totalPages} onPageChange={setCurrentPage} />
    </main>
  )
}

