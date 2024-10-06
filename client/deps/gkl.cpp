#include <functional>
#include <windows.h>
#include <iostream>
#include <unordered_set>

#define TOGGLE_HK 1001
#define EXIT_HK 1002

using namespace std;

HHOOK hHook;

WPARAM KEY_PRESS = 256;
WPARAM KEY_RELEASE = 257;

unordered_set<DWORD> passthroughKeys = {VK_F13, VK_F14, VK_F15, VK_F16, VK_F17, VK_F18, VK_F19, VK_F20, VK_F21, VK_SCROLL};

bool remoteControlActive = false;

LONG_PTR supressKeyStroke = 1;

void ToggleScrollLock(BOOL shouldDisable)
{
    if (shouldDisable && (GetKeyState(VK_SCROLL) & 0x0001) == 0)
        return;

    INPUT input[2];
    ::ZeroMemory(input, sizeof(input));
    input[0].type = input[1].type = INPUT_KEYBOARD;
    input[0].ki.wVk = input[1].ki.wVk = VK_SCROLL;
    input[1].ki.dwFlags = KEYEVENTF_KEYUP;
    ::SendInput(2, input, sizeof(INPUT));
}

LRESULT CALLBACK LowLevelKeyboardProc(int nCode, WPARAM wParam, LPARAM lParam)
{
    if (nCode == HC_ACTION)
    {
        KBDLLHOOKSTRUCT *kbStruct = (KBDLLHOOKSTRUCT *)(lParam);
        DWORD vkCode = reinterpret_cast<PKBDLLHOOKSTRUCT>(lParam)->vkCode;

        cout << vkCode << ":" << wParam << endl;

        if (passthroughKeys.count(vkCode) > 0)
        {
            return CallNextHookEx(hHook, nCode, wParam, lParam);
        }
    }

    return supressKeyStroke;
}

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow)
{
    ToggleScrollLock(true);

    RegisterHotKey(NULL, TOGGLE_HK, 0, VK_F14);
    RegisterHotKey(NULL, EXIT_HK, 0, VK_F15);

    MSG msg;

    while (GetMessage(&msg, NULL, 0, 0))
    {
        if (msg.message == WM_HOTKEY && msg.wParam == EXIT_HK)
            break;

        if (msg.message == WM_HOTKEY && msg.wParam == TOGGLE_HK)
        {
            remoteControlActive = !remoteControlActive;

            ToggleScrollLock(false);

            if (!remoteControlActive)
            {
                WinExec("C:\\Users\\aron koffler\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\nuphy2.ahk", 0);
                UnhookWindowsHookEx(hHook);
            }
            else
            {
                hHook = SetWindowsHookEx(WH_KEYBOARD_LL, LowLevelKeyboardProc, hInstance, 0);
            }
        }
    }

    ToggleScrollLock(true);

    return 0;
}