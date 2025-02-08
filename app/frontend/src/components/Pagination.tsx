import KeyboardDoubleArrowRightIcon from '@mui/icons-material/KeyboardDoubleArrowRight';
import KeyboardDoubleArrowLeftIcon from '@mui/icons-material/KeyboardDoubleArrowLeft';
import KeyboardArrowRightIcon from '@mui/icons-material/KeyboardArrowRight';
import KeyboardArrowLeftIcon from '@mui/icons-material/KeyboardArrowLeft';

interface PaginationProps {
  currentPage: number
  totalPages: number
  onPageChange: (page: number) => void
}

export default function Pagination({ currentPage, totalPages, onPageChange }: PaginationProps) {
  const getPageNumbers = () => {
    if (totalPages <= 3) {
      return Array.from({ length: totalPages }, (_, i) => i + 1)
    }

    if (currentPage <= 2) return [1, 2, 3]
    if (currentPage >= totalPages - 1) return [totalPages - 2, totalPages - 1, totalPages]

    return [currentPage - 1, currentPage, currentPage + 1]
  }

  return (
    <div className="flex justify-center items-center space-x-2 mt-8">
      <button
        onClick={() => onPageChange(1)}
        className="p-2 rounded-md bg-white text-primary hover:bg-primary hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <KeyboardDoubleArrowLeftIcon className="w-5 h-5" />
      </button>

      <button
        onClick={() => onPageChange(Math.max(1, currentPage - 1))}
        disabled={currentPage === 1}
        className="p-2 rounded-md bg-white text-primary hover:bg-primary hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <KeyboardArrowLeftIcon className="w-5 h-5" />
      </button>

      {getPageNumbers().map((pageNumber) => (
        <button
          key={pageNumber}
          onClick={() => onPageChange(pageNumber)}
          className={`px-4 py-2 rounded-md ${
            currentPage === pageNumber
              ? "bg-primary text-white"
              : "bg-white text-primary hover:bg-primary hover:text-white"
          } transition-colors duration-300`}
        >
          {pageNumber}
        </button>
      ))}

      <button
        onClick={() => onPageChange(Math.min(totalPages, currentPage + 1))}
        disabled={currentPage === totalPages}
        className="p-2 rounded-md bg-white text-primary hover:bg-primary hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <KeyboardArrowRightIcon className="w-5 h-5" />
      </button>

      <button
        onClick={() => onPageChange(totalPages)}
        className="p-2 rounded-md bg-white text-primary hover:bg-primary hover:text-white transition-colors duration-300 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <KeyboardDoubleArrowRightIcon className="w-5 h-5" />
      </button>
    </div>
  )
}

