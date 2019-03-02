import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns

df = pd.read_excel("../testdata/plot.xlsx")
plt.rcParams['font.sans-serif'] = ['SimHei']  # 步骤一（替换sans-serif字体）
plt.rcParams['axes.unicode_minus'] = False  # 步骤二（解决坐标轴负数的负号显示问题）


# 直方图
def zhifangtu():
    fig = plt.figure()
    ax = fig.add_subplot(1, 1, 1)
    ax.hist(df['Age'], bins=7)

    plt.title('Age distribution')
    plt.xlabel('Age')
    plt.ylabel('#Employee')
    plt.show()


# 箱线图
def boxplot():
    fig = plt.figure()
    ax = fig.add_subplot(1, 1, 1)
    ax.boxplot(df['Age'])
    plt.show()


def violin():
    sns.violinplot(df['Age'], df['Gender'])
    sns.despine()


if __name__ == "__main__":
    # zhifangtu()
    # boxplot()
    violin()
