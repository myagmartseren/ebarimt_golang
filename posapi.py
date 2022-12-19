import ctypes
import json
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


if __name__ == "__main__":
    temp = PosApi()
    print("checkApi:\t", json.loads(temp.checkApi()))
    print("getInformation:\t", json.loads(temp.getInformation()))

    print("call function:\t", json.loads(temp.callFunction("regNo","ЙЮ01312715")))
    test = "000000004012345678901234567890"
    print(test,len(test))
    returnbill_input = json.dumps({
        "returnBillId": test,
        "date": "2022-12-13 11:22:22"})
    result = temp.returnBill(returnbill_input.encode("ascii"))
    print("return bill:\t", json.loads(result))
    result = temp.sendData()
    print("sendata:\t", json.loads(result))
    req = temp.put(json.dumps({
        "amount": "100000",
        "vat": "100",
        "cashAmount": "100000",
        "nonCashAmount": "100000",
        "cityTax": "",
        "districtCode": "1",
        "posNo": "12345",
        "returnBillId": "111111111111111111111111111111111",
        "invoiceId": "111111111111111111111111111111111",
        "reportMonth": "2022-12",
        "branchNo": "333",
        "stocks": [
            {
                "code": "test ym aa",
                "name": "test123",
                "measureUnit": "unit",
                "qty": "12",
                "unitPrice": "12",
                "totalAmount": "12",
                "cityTax": "12",
                "vat": "12",
                "barCode": "235"
            }
        ],
        "bankTransactions": [
            {
                "rrn": "111111111111",
                "bankId": "01",
                "terminalId": "asd123",
                "approvalCode": "123456abcdef",
                "amount": "12"
            }
        ]}).encode("ascii"))

    print("put:\t\t", json.loads(req))
