
# 传入git地址和git凭据，获取所有分支
from git.repo import Repo
import os, shutil
from getpass import getpass

git_repo = "http://root:1234.com@192.168.1.90:88/root/java-demo.git"
git_credential = ""
clone_dir = os.path.join(os.getcwd(), 'test')

if os.path.exists(clone_dir):
    shutil.rmtree(clone_dir)

Repo.clone_from(git_repo, to_path=clone_dir)
# 获取分支
repo = Repo(clone_dir)
branch = []
for ref in repo.remote().refs:
    if ref.remote_head != 'HEAD':
        branch.append(ref.remote_head)

