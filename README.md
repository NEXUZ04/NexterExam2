# NexterExam2

Working Flow
-------------
[Initial all cassets]
        V
[Show warning to customer which cash type are full]
        V
[Customer insert product price]
        V
[Customer insert cash note/coin one by one] 
        V
[Deposit cash to casset]
    V            V
    V   [If target casset is full: return customer cash back and let them insert again]
    V
[Insert cash until more than product price]
    V
[Machine withdraw cash type accendingly for change money by each casset]
    V
[If complete/not enough for that cash type, Machine will find cash in next casset]
    V            V
    V   [If cannot find change money enough from all casset: return all inserted cash and fed cash back to customer and casset]
    V
[If complete, send change money to customer]


Architect
----------
Package core: DB and api for handling DB
         |
         V
Package api: api for handling transction
         |
         V
Package main: Machine interface

