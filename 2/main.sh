
[ "$1" = "" ] && echo "USAGE: $0 <filename>" && exit

P1=0
P2=0

while read -r line
do
    min=`echo $line | cut -d'-' -f1`
    max=`echo $line | cut -d'-' -f2 | awk '{print $1}'`
    letter=`echo $line | cut -d':' -f1 | awk '{print $2}'`
    password=`echo $line | awk '{print $3}'`
    #count=$(echo $password | awk -F"$letter" '{print NF-1}')
    count=`echo $password | tr -c -d $letter | wc -c`
    if [ "$count" -lt $max ] || [ "$count" -gt $min ]; then
        ((P1+=1))
    else
        :
    fi
    #echo $line
done < $1

echo $P1
echo $P2