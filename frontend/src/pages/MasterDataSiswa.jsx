import React, { useState, useEffect } from "react";
import {
  useReactTable,
  getCoreRowModel,
  getSortedRowModel,
  getPaginationRowModel,
  getFilteredRowModel,
  flexRender,
  filterFns,
} from "@tanstack/react-table";

const MasterDataSiswa = () => {
  const [filtering, setFiltering] = useState("");
  const [data, setData] = useState([]);

  // Fungsi Ambil Data dari API
  const fetchData = async () => {
    try {
      const token = localStorage.getItem("token"); // Ambil token dari localStorage
      const response = await fetch("http://localhost:8096/api/listsiswa", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) throw new Error("Gagal mengambil data");

      const result = await response.json();
      setData(result.data); // Sesuaikan dengan struktur data API
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchData();
    const connectWebSocket = () => {
      const socket = new WebSocket("ws://localhost:8096/ws/listsiswa");
  
      socket.onopen = () => console.log("WebSocket Connected!");
      socket.onmessage = (event) => {
        const updatedData = JSON.parse(event.data);
        setData(updatedData);
      };
      socket.onerror = (error) => console.error("WebSocket Error:", error);
      socket.onclose = () => {
        console.log("WebSocket Disconnected! Reconnecting...");
        setTimeout(connectWebSocket, 5000);
      };
    };
  
    connectWebSocket();

  }, []);

  // Fungsi Edit
  const handleEdit = (row) => {
    console.log("Edit data:", row);
    alert(`Edit siswa: ${row.name}`);
  };

  // Fungsi Delete
  const handleDelete = (id) => {
    const newData = data.filter((row) => row.id !== id);
    setData(newData);
  };

  const columns = [
    { accessorKey: "id", header: "No" },
    { accessorKey: "name", header: "Nama" },
    { accessorKey: "nik", header: "NIK" },
    { accessorKey: "email", header: "Email" },
    { accessorKey: "alamat", header: "Alamat" },
    { accessorKey: "created_at", header: "Created At" },
    {
      accessorKey: "aksi",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="flex gap-1">
          <button
            className="bg-blue-500 text-white px-1 py-0.5 text-sm rounded"
            onClick={() => handleEdit(row.original)}
          >
            Edit
          </button>
          <button
            className="bg-red-500 text-white px-1 py-0.5 text-sm rounded"
            onClick={() => handleDelete(row.original.id)}
          >
            Del
          </button>
        </div>
      ),
    },
  ];

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    globalFilterFn: filterFns.includesString,
    state: {
      globalFilter: filtering,
    },
  });

  return (
    <div className="p-6">
      <h1 className="text-2xl md:text-3xl text-gray-800 dark:text-gray-100 font-bold">
        Master Data Siswa
      </h1>

      {/* Search Input */}
      <input
        type="text"
        className="mb-4 p-2 bg-transparent border-b border-gray-300 text-gray-800 dark:text-gray-100 focus:outline-none focus:border rounded w-full"
        placeholder="Cari siswa..."
        value={filtering}
        onChange={(e) => setFiltering(e.target.value)}
      />

      {/* Table */}
      <table className="px-4 sm:px-6 lg:px-8 py-8 w-full max-w-9xl mx-auto">
        <thead className="text-gray-800 dark:text-gray-100 font-bold">
          {table.getHeaderGroups().map((headerGroup) => (
            <tr key={headerGroup.id}>
              {headerGroup.headers.map((header) => (
                <th
                  key={header.id}
                  className="px-4 py-2 border text-left cursor-pointer"
                  onClick={header.column.getToggleSortingHandler()}
                >
                  {flexRender(header.column.columnDef.header, header.getContext())}
                  {header.column.getIsSorted() === "asc" ? " 🔼" : header.column.getIsSorted() === "desc" ? " 🔽" : ""}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody>
          {table.getRowModel().rows.map((row) => (
            <tr key={row.id} className="hover:bg-gray-200">
              {row.getVisibleCells().map((cell) => (
                <td key={cell.id} className="px-4 py-2 border">
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>

      {/* Pagination */}
      <div className="flex justify-between mt-4">
        <button
          onClick={() => table.previousPage()}
          disabled={!table.getCanPreviousPage()}
          className="text-gray-800 dark:text-gray-100 font-bold"
        >
          Prev
        </button>
        <span>
          Page {table.getState().pagination.pageIndex + 1} of {table.getPageCount()}
        </span>
        <button
          onClick={() => table.nextPage()}
          disabled={!table.getCanNextPage()}
          className="text-gray-800 dark:text-gray-100 font-bold"
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default MasterDataSiswa;
