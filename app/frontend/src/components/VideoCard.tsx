import Link from "next/link"

interface VideoCardProps {
  title: string
  thumbnailUrl: string
  videoUrl: string
}

export default function VideoCard({ title, thumbnailUrl, videoUrl }: VideoCardProps) {
  return (
    <Link href={videoUrl} className="group">
      <div className="bg-white shadow-md overflow-hidden transition-transform duration-300 transform hover:scale-105 relative">
        <div className="relative h-48 justify-center flex">
          <img src={thumbnailUrl} alt={title} />
        </div>
        <div className="p-4">
          <h3 className="text-lg font-semibold text-text group-hover:text-primary transition-colors duration-300 line-clamp-2 h-20">
            {title}
          </h3>
        </div>
      </div>
    </Link>
  )
}