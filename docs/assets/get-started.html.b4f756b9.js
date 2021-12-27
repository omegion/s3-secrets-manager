import{r as a,o as r,a as l,b as n,d as t,F as o,e as s,c as i}from"./app.3e9e7185.js";import{_ as c}from"./plugin-vue_export-helper.21dcd24c.js";const p={},d=n("h1",{id:"get-started",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#get-started","aria-hidden":"true"},"#"),s(" Get Started")],-1),b=n("h2",{id:"prerequisites",tabindex:"-1"},[n("a",{class:"header-anchor",href:"#prerequisites","aria-hidden":"true"},"#"),s(" Prerequisites")],-1),u=n("p",null,[s("Bitwarden SSH Manager uses "),n("code",null,"bw"),s(" CLI command in background. So, you will need to install it and login to be able to use the SSH Manager.")],-1),m={href:"https://bitwarden.com/help/article/cli/#quick-start",target:"_blank",rel:"noopener noreferrer"},h=s("Bitwarden CLI"),g=i(`<h2 id="install" tabindex="-1"><a class="header-anchor" href="#install" aria-hidden="true">#</a> Install</h2><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>go get -u github.com/omegion/bw-ssh
</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>This will install <code>bw-ssh</code> binary to your <code>GOPATH</code>.</p><p>Let&#39;s verify that the binary has installed successfully.</p><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>\u276F bw-ssh --help            
CLI <span class="token builtin class-name">command</span> to manage SSH keys stored on Bitwarden

Usage:
  bw-ssh <span class="token punctuation">[</span>command<span class="token punctuation">]</span>

Available Commands:
  <span class="token function">add</span>         Add SSH key to Bitwarden.
  get         Get SSH key from Bitwarden.
  <span class="token builtin class-name">help</span>        Help about any <span class="token builtin class-name">command</span>
  version     Print the version/build number

Flags:
  -h, --help   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> bw-ssh

Use <span class="token string">&quot;bw-ssh [command] --help&quot;</span> <span class="token keyword">for</span> <span class="token function">more</span> information about a command.
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br></div></div>`,5);function _(k,f){const e=a("ExternalLinkIcon");return r(),l(o,null,[d,b,u,n("ul",null,[n("li",null,[n("a",m,[h,t(e)])])]),g],64)}var S=c(p,[["render",_]]);export{S as default};
