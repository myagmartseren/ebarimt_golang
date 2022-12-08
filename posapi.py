import ctypes

DEFAULT_LIB = "libPosAPI"
LIB_DIR = "/usr/lib/{}.so"


class PosApi:
    def __init__(self, lib=DEFAULT_LIB):
        if lib is None:
            libdir = "/usr/lib/libPosAPI.so"
        else:
            libdir = LIB_DIR.format(lib)
        self.API = ctypes.cdll.LoadLibrary(libdir)

        self.API.checkApi.restype = ctypes.c_char_p
        self.API.getInformation.restype = ctypes.c_char_p
        self.API.callFunction.restype = ctypes.c_char_p
        self.API.put.restype = ctypes.c_char_p
        self.API.returnBill.restype = ctypes.c_char_p
        self.API.sendData.restype = ctypes.c_char_p

    def checkApi(self):
        res_check_api = self.API.checkApi()
        return res_check_api

    def getInformation(self):
        res_get_information = self.API.getInformation()
        return res_get_information

    def callFunction(self, functionName, params):
        response_check_api = self.API.callFunction(functionName, params)
        return response_check_api

    def put(self, params):
        res = self.API.put(params)
        return res

    def returnBill(self, params):
        res = self.API.returnBill(params)
        return res

    def sendData(self):
        res = self.API.sendData()
        return res
