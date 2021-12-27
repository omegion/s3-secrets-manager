import{c as n}from"./app.7023e78d.js";import{_ as s}from"./plugin-vue_export-helper.21dcd24c.js";const a={},e=n(`<h1 id="introduction" tabindex="-1"><a class="header-anchor" href="#introduction" aria-hidden="true">#</a> Introduction</h1><p>AWS S3 is widely used storage system secured very well. S3 Secrets Manager uses S3 to store your secrets in a structured way that you can access anywhere any time.</p><p>You can set your S3 bucket to store your objects with KSM key to add extra security. In addition, you can create IAM role to limit the tools operations.</p><div class="language-bash ext-sh line-numbers-mode"><pre class="language-bash"><code>S3 Secrets Management <span class="token keyword">for</span> AWS S3.

Usage:
  s3sm <span class="token punctuation">[</span>command<span class="token punctuation">]</span>

Available Commands:
  completion  Generate the autocompletion script <span class="token keyword">for</span> the specified shell
  <span class="token builtin class-name">help</span>        Help about any <span class="token builtin class-name">command</span>
  secret      Secret operations.
  version     Print the version/build number

Flags:
      --config string      config <span class="token function">file</span> <span class="token punctuation">(</span>default is <span class="token environment constant">$HOME</span>/.s3sm/config.yaml<span class="token punctuation">)</span>
  -h, --help               <span class="token builtin class-name">help</span> <span class="token keyword">for</span> s3sm
      --interactive        Set the interactivity <span class="token punctuation">(</span>default <span class="token boolean">true</span><span class="token punctuation">)</span>
      --logFormat string   Set the logging format. One of: text<span class="token operator">|</span>json <span class="token punctuation">(</span>default <span class="token string">&quot;text&quot;</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>default <span class="token string">&quot;text&quot;</span><span class="token punctuation">)</span>
      --logLevel string    Set the logging level. One of: debug<span class="token operator">|</span>info<span class="token operator">|</span>warn<span class="token operator">|</span>error <span class="token punctuation">(</span>default <span class="token string">&quot;info&quot;</span><span class="token punctuation">)</span>

Use <span class="token string">&quot;s3sm [command] --help&quot;</span> <span class="token keyword">for</span> <span class="token function">more</span> information about a command.
</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br></div></div>`,4);function t(o,p){return e}var l=s(a,[["render",t]]);export{l as default};
