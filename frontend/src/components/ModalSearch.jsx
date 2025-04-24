import React, { useRef, useEffect } from 'react';
import Transition from '../utils/Transition';

function LogoutModal({ modalOpen, setModalOpen, onLogout }) {
  const modalContent = useRef(null);

  // Close modal when clicking outside
  useEffect(() => {
    const clickHandler = ({ target }) => {
      if (!modalOpen || modalContent.current.contains(target)) return;
      setModalOpen(false);
    };
    document.addEventListener('click', clickHandler);
    return () => document.removeEventListener('click', clickHandler);
  });

  // Close modal when pressing ESC
  useEffect(() => {
    const keyHandler = ({ keyCode }) => {
      if (!modalOpen || keyCode !== 27) return;
      setModalOpen(false);
    };
    document.addEventListener('keydown', keyHandler);
    return () => document.removeEventListener('keydown', keyHandler);
  });

  return (
    <>
      {/* Modal backdrop */}
      <Transition
        className="fixed inset-0 bg-gray-900/30 z-50 transition-opacity"
        show={modalOpen}
        enter="transition ease-out duration-200"
        enterStart="opacity-0"
        enterEnd="opacity-100"
        leave="transition ease-out duration-100"
        leaveStart="opacity-100"
        leaveEnd="opacity-0"
        aria-hidden="true"
      />
      {/* Modal dialog */}
      <Transition
        className="fixed inset-0 z-50 flex items-center justify-center px-4 sm:px-6"
        role="dialog"
        aria-modal="true"
        show={modalOpen}
        enter="transition ease-in-out duration-200"
        enterStart="opacity-0 translate-y-4"
        enterEnd="opacity-100 translate-y-0"
        leave="transition ease-in-out duration-200"
        leaveStart="opacity-100 translate-y-0"
        leaveEnd="opacity-0 translate-y-4"
      >
        <div ref={modalContent} className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-full max-w-sm">
          <h3 className="text-lg font-semibold text-gray-800 dark:text-gray-100">Konfirmasi Logout</h3>
          <p className="text-gray-600 dark:text-gray-400 mt-2">Apakah Anda yakin ingin keluar?</p>
          <div className="flex justify-end mt-4">
            <button
              className="px-4 py-2 mr-2 bg-gray-300 dark:bg-gray-700 rounded-lg hover:bg-gray-400 dark:hover:bg-gray-600"
              onClick={() => setModalOpen(false)}
            >
              Batal
            </button>
            <button
              className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
              onClick={onLogout}
            >
              Logout
            </button>
          </div>
        </div>
      </Transition>
    </>
  );
}

export default LogoutModal;
