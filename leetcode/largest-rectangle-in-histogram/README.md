
# O(n^2) solution:
ö�����п�����ʼ/��ֹλ�ã���һ�������
```go
for i := 0; i < n; i++ {
    minh := h[i]
    for j := i+1; j <= n; j++ {
        minh = min(minh, h[j-1])
        if (j-i) * min > maxArea {
            maxArea = (j-i)*min
        }
    }
}
```

# insight 

�����Ŀ�Ĺؼ� insight: ���մ�һ���������е�ĳ���߶� h[i] ����������չ���ɵľ��Ρ�


## ֤��
ʹ�÷�֤��֤�������� **�γ����մ𰸵ľ��Σ���߶Ȳ��������������еĸ߶�**, �Ǹþ������Ϊ A1��
��һ������ͨ�����Ӹþ��θ߶ȵķ�ʽ����������� A2����Ϊ A2 > A1���ʼ��費������


# solution

## ����һ
��ֱ�۵��뷨������ÿ�� h[i]���ֱ��ҵ������/�ұ������ < h[i] ��λ�ã��� `h[i] * (lessFromRigit[i] - lessFromLeft[i] - 1)` ���Ը��µ�ǰ�� maxArea���������� h[i] �󣬼���ô𰸡�

���㷨����������Ȼ�� O(n^2)����Ϊ��ҪΪÿ�� h[i] ���� lessFromLeft[i]��

```go
for i := 1; i < n; i++ {
    j = i-1
    for j >= 0 && h[j] >= h[i] {
        j--
    }
    lessFromLeft[i] = j
}
```

### ���㷨�� O(n^2) ���͵� O(n) �ĵ㾦֮��

��һ���� h[i] ����ڵ㲢����������Ϊû�г�������Ѿ�������� `lessFromLeft[j], 0 <= j < i`��
��Ȼ `h[j] >= h[i]`����Ȼ���ǿ����޷�ֱ���ҵ���һ���� h[i] С����ڵ㣨�� lessFromLeft[i]�������ǿ������ҵ�
��һ���� h[j] С����ڵ㣬ֱ���������� >= h[j] �Ľڵ㡣

�� h[j] С�Ľڵ������ֿ���
1. С�� h[i]�����ҵ� lessFromLeft[i]
2. ��Ȼ >= h[i]������Ѱ�ұȸýڵ�С����ڵ㣬ֱ������ 1


```go
for i := 1; i < n; i++ {
    j = i-1
    for j >= 0 && h[j] >= h[i] {
        j = lessFromLeft[j]
    }
    lessFromLeft[i] = j
}
```

### ���֤���ڲ�ѭ���ľ�̯���Ӷ��� O(n) ��?

֤�������� `0 <= i < n`, lessFromLeft[i] ���ڴ�ѭ�������ᱻ����һ��

����򵥵������ʼ���� h ���е��������������� `h[j] < h[i], j < i`���޷���������ڲ�ѭ�������� `h[j] >= h[i]`���� lessFromLeft[i] �����ʴ���Ϊ 0��

�� h �г��ַ�����������ʱ��lessFromLeft[i] ���п��ܱ����ʵ���

�� [1, 2, 3, 4, 5, 3, 8, 2] ����, i = 4 ������:

i        | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 
-------- | - | - | - | - | - | - | - | -
h        | 1 | 2 | 3 | 4 | 5 | 3 | 8 | 2
less     |-1 | 0 | 1 | 2 | 3 | N/A | N/A | N/A
���ʴ��� | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0

i = 5 ������

i        | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 
-------- | - | - | - | - | - | - | - | -
h        | 1 | 2 | 3 | 4 | 5 | 3 | 8 | 2
less     |-1 | 0 | 1 | 2 | 3 | 1 | N/A | N/A
���ʴ��� | 0 | 0 | 1 | 1 | 1 | 0 | 0 | 0

���Կ�������Ϊ h[5] = 3���״γ����˷��������У�lessFromLeft[j] (`2 <= j <= 4`) ���ڲ�ѭ���б�����������Ϊ `h[j] >= h[i]`����ô����֮��ı����У��Ƿ񻹻��ٴα��� lessFromLeft[j] �أ�

���ڶ��γ��ַ���������ʱ��i = 7 ������:

i        | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 
-------- | - | - | - | - | - | - | - | -
h        | 1 | 2 | 3 | 4 | 5 | 3 | 8 | 2
less     |-1 | 0 | 1 | 2 | 3 | 1 | 5 | 0
���ʴ��� | 0 | 1 | 1 | 1 | 1 | 1 | 1 | 0

���Կ�������� lessFromLeft[7] ���ڲ�ѭ���ı���˳���ǣ�`6->5->1->0`����Ϊ lessFromLeft[5] �Ĵ��ڣ������˶� `2 <= j <= 4` �ı�����

�ܽ�һ�£�
1. ���� h �ĵ��������У���������κ� lessFromLeft[i]
2. �ɱ����ʵ��� lessFromLeft[j] ʼ�ձ����ŵ�������������
3. �� `h[j] >= h[i] && i > j` ʱ��lessFromLeft[i] �γ��˶� [j, i) �����ϣ�ʹ�䲻���ٱ�����
4. lessFromLeft[j] һ�������ʣ�˵�� `h[j] >= h[i] && i > j`����� 3��lessFromLeft[j] �����ٱ�����

## ������

����һ�ĸ��Ӷȷ����������Ƶ����������ṩ������ֱ���ĸ��Ӷȷ���������

��������� insight: ���մ�һ���������е�ĳ���߶� h[i]��

����һ������ h[i]������ lessFromLeft[i] �� lessFromRight[i]

���������� `h[i+1] < h[i]` ʱ��h[i] �� lessFromRight ��Ϊ `i+1`��lessFromLeft �����ڱ��� h �Ĺ����м�¼������ά��һ���߶ȵ����ǽ���λ������ Q������֮�⣬h[i+1] ������ h[i] �� lessFromRight���������� Q ������λ�õ� lessFromRight��ֻҪ���� `h[Q[len(Q)-1]] > h[i+1]`

���Ӷ�֤����һ��ѭ������ h��ÿ��λ�������� Q һ�Σ����Ӷ� O(n)